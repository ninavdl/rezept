package api

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sour-dough/rezept-backend/db"
)

type Ingredient struct {
	Name   string          `validate:"required"`
	Amount decimal.Decimal `validate:"numeric"`
	Unit   string
	Note   string `validate:"max=10"`
}

func (i *Ingredient) toDB() db.Ingredient {
	return db.Ingredient{
		Name:   i.Name,
		Amount: i.Amount,
		Unit:   i.Unit,
		Note:   i.Note,
	}
}

func newIngredient(i *db.Ingredient) Ingredient {
	return Ingredient{
		Name:   i.Name,
		Amount: i.Amount,
		Unit:   i.Unit,
		Note:   i.Note,
	}
}

type Step struct {
	Text  string `validate:"required"`
	Image *Image
}

func (s *Step) toDB() db.Step {
	var imageID uint = 0
	if s.Image != nil {
		imageID = s.Image.ID
	}
	return db.Step{
		Text:    s.Text,
		ImageID: imageID,
	}
}

func (api *API) newStep(s *db.Step) Step {
	var img *Image
	if s.Image != nil {
		image := api.newImage(s.Image)
		img = &image
	}
	return Step{
		Text:  s.Text,
		Image: img,
	}
}

type Recipe struct {
	ID               uint
	Name             string `validate:"required"`
	Description      string
	ShortDescription string
	Servings         uint         `validate:"required"`
	Ingredients      []Ingredient `validate:"dive"`
	Steps            []Step       `validate:"dive"`
	Tags             []string
	Image            *Image
	Creator          *User
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Published        bool
}

func (r *Recipe) toDB() db.Recipe {
	steps := make([]db.Step, len(r.Steps))
	ingredients := make([]db.Ingredient, len(r.Ingredients))
	tags := make([]db.Tag, len(r.Tags))
	for i, s := range r.Steps {
		steps[i] = s.toDB()
	}
	for i, in := range r.Ingredients {
		ingredients[i] = in.toDB()
	}
	for i, t := range r.Tags {
		tags[i].Tag = t
	}
	var creator *db.User
	if r.Creator != nil {
		c := r.Creator.toDB()
		creator = &c
	}
	var imgID uint = 0
	if r.Image != nil {
		imgID = r.Image.ID
	}
	return db.Recipe{
		RecipeInfo: db.RecipeInfo{
			ID:               r.ID,
			Name:             r.Name,
			ShortDescription: r.ShortDescription,
			CreatedAt:        r.CreatedAt,
			UpdatedAt:        r.UpdatedAt,
			ImageID:          imgID,
			Published:        r.Published,
		},
		Creator:     creator,
		Description: r.Description,
		Ingredients: ingredients,
		Steps:       steps,
		Servings:    r.Servings,
		Tags:        tags,
	}
}

func (api *API) newRecipe(r *db.Recipe) Recipe {
	steps := make([]Step, len(r.Steps))
	ingredients := make([]Ingredient, len(r.Ingredients))
	tags := make([]string, len(r.Tags))
	for i, s := range r.Steps {
		steps[i] = api.newStep(&s)
	}
	for i, in := range r.Ingredients {
		ingredients[i] = newIngredient(&in)
	}
	for i, t := range r.Tags {
		tags[i] = t.Tag
	}
	var creator *User
	if r.Creator != nil {
		c := newUser(r.Creator)
		creator = &c
	}
	var image *Image
	if r.Image != nil {
		i := api.newImage(r.Image)
		image = &i
	}
	return Recipe{
		ID:               r.ID,
		Name:             r.Name,
		Description:      r.Description,
		ShortDescription: r.ShortDescription,
		Servings:         r.Servings,
		Ingredients:      ingredients,
		Steps:            steps,
		Tags:             tags,
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
		Creator:          creator,
		Image:            image,
		Published:        r.Published,
	}
}

func (api *API) getRecipe(r request) error {
	id, err := strconv.Atoi(r.params.ByName("id"))
	if err != nil {
		return err
	}

	dbRecipe := api.db.GetRecipe(uint(id))
	if dbRecipe == nil {
		return r.writeError("Recipe not found", 404)
	}

	recipe := api.newRecipe(dbRecipe)

	if !recipe.Published && r.user != nil && !r.user.CanModifyRecipe(&recipe) {
		return r.writeError("User cannot access recipe", 403)
	}

	return r.writeJson(recipe)
}

func (api *API) putRecipe(r request) error {
	var recipe Recipe
	err := json.NewDecoder(r.req.Body).Decode(&recipe)
	if err != nil {
		return err
	}

	if !r.validateStruct(&recipe) {
		return nil
	}

	recipe.Creator = r.user

	dbRecipe := recipe.toDB()
	err = api.db.PutRecipe(&dbRecipe)
	if err != nil {
		return err
	}

	r.code = 201
	return r.writeJson(api.newRecipe(&dbRecipe))
}

func (api *API) updateRecipe(r request) error {
	id, err := strconv.Atoi(r.params.ByName("id"))
	if err != nil {
		return err
	}

	var recipe Recipe
	err = json.NewDecoder(r.req.Body).Decode(&recipe)
	if err != nil {
		return err
	}

	if recipe.ID != 0 && recipe.ID != uint(id) {
		return r.writeError("Cannot update ID of recipe", 404)
	}

	originalRecipeDB := api.db.GetRecipe(recipe.ID)
	if originalRecipeDB == nil {
		return r.writeError("Recipe does not exist", 404)
	}

	originalRecipe := api.newRecipe(originalRecipeDB)
	if !r.user.CanModifyRecipe(&originalRecipe) {
		return r.writeError("Forbidden", 403)
	}

	recipe.Creator = originalRecipe.Creator

	dbRecipe := recipe.toDB()
	err = api.db.UpdateRecipe(&dbRecipe)
	if err != nil {
		return err
	}

	go func() {
		err := api.deleteUnlinkedImages()
		if err != nil {
			log.Println(err)
		}
	}()

	return r.writeJson(api.newRecipe(&dbRecipe))
}

func (api *API) deleteRecipe(r request) error {
	id, err := strconv.Atoi(r.params.ByName("id"))
	if err != nil {
		return err
	}

	recipeDb := api.db.GetRecipe(uint(id))
	if recipeDb == nil {
		return r.writeError("Recipe not found", 404)
	}

	recipe := api.newRecipe(recipeDb)

	if !r.user.CanModifyRecipe(&recipe) {
		return r.writeError("Forbidden", 403)
	}

	err = api.db.DeleteRecipe(id)
	if err != nil {
		return err
	}

	go func() {
		err := api.deleteUnlinkedImages()
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
