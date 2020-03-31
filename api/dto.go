package api

import (
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

type RecipeInfo struct {
	ID               uint
	Name             string
	ShortDescription string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Image            *Image
}

func (r *RecipeInfo) toDB() db.RecipeInfo {
	ri := db.RecipeInfo{
		ID:               r.ID,
		Name:             r.Name,
		ShortDescription: r.ShortDescription,
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
	}
	if r.Image != nil {
		ri.ImageID = r.Image.ID
	}
	return ri
}

func (api *API) newRecipeInfo(r *db.RecipeInfo) RecipeInfo {
	var img *Image
	if r.Image != nil {
		apiImage := api.newImage(r.Image)
		img = &apiImage
	}
	return RecipeInfo{
		ID:               r.ID,
		Name:             r.Name,
		ShortDescription: r.ShortDescription,
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
		Image:            img,
	}
}

type RecipeList struct {
	Recipes []RecipeInfo
	Page    uint
	Pages   uint
	Results uint
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
	}
}

type User struct {
	ID          uint
	Username    string
	DisplayName string
	IsAdmin     bool
}

func (u *User) toDB() db.User {
	return db.User{
		ID:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		IsAdmin:     u.IsAdmin,
	}
}

func newUser(u *db.User) User {
	return User{
		ID:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		IsAdmin:     u.IsAdmin,
	}
}

type UserRegistration struct {
	Username    string `validate:"required"`
	DisplayName string
	Password    string `validate:"required,min=8"`
}

type LoginRequest struct {
	Username string
	Password string
}

type Image struct {
	ID           uint
	URL          string
	ThumbnailURL string
}

func (api *API) newImage(img *db.Image) Image {
	return Image{
		ID:           img.ID,
		URL:          api.GetImageURL(img.ID),
		ThumbnailURL: api.GetThumbnailURL(img.ID),
	}
}

func (api *API) imageToDB(img *Image) db.Image {
	return db.Image{
		ID: img.ID,
	}
}
