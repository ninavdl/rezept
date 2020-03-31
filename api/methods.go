package api

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/sour-dough/rezept-backend/db"
	"golang.org/x/crypto/bcrypt"
)

func userCanModifyRecipe(user *User, recipe *Recipe) bool {
	if user == nil {
		return false
	}

	if user.IsAdmin {
		return true
	}

	return recipe != nil && recipe.Creator != nil && recipe.Creator.ID == user.ID
}

func (api *API) listRecipes(r request) error {
	queryValues := r.req.URL.Query()
	pageVal := queryValues.Get("page")
	var page int
	if pageVal == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pageVal)
		if err != nil || page < 1 {
			return r.writeError("page must be integer > 0", 400)
		}
	}

	fmt.Println(pageVal, page)

	tags := queryValues["tag"]
	username := queryValues.Get("user")
	keywords := queryValues["keyword"]

	var userID uint
	if username != "" {
		user := api.db.GetUserByName(username)
		if user == nil {
			return r.writeJson(RecipeList{Recipes: make([]RecipeInfo, 0), Page: 0, Pages: 0, Results: 0})
		}
		userID = user.ID
	}

	if tags != nil && len(tags) > 10 {
		return r.writeError("Cannot search for more than 10 tags at once", 400)
	}

	if keywords != nil && len(keywords) > 10 {
		return r.writeError("Cannot search for more than 10 keywords at once", 400)
	}

	filter := db.RecipeFilter{
		Tags:     tags,
		Keywords: keywords,
		User:     userID,
	}

	const limit = 10
	offset := (page - 1) * limit
	count := api.db.CountRecipes(filter)

	dbRecipes, err := api.db.GetRecipes(filter, uint(offset), limit)
	if err != nil {
		return err
	}

	pages := uint(math.Ceil(float64(count) / float64(limit)))

	recipes := make([]RecipeInfo, len(dbRecipes))
	for i, r := range dbRecipes {
		recipes[i] = api.newRecipeInfo(&r)
	}

	return r.writeJson(RecipeList{
		Recipes: recipes,
		Page:    uint(page),
		Pages:   pages,
		Results: count,
	})
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
	return r.writeJson(api.newRecipe(dbRecipe))
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
	if !userCanModifyRecipe(r.user, &originalRecipe) {
		return r.writeError("Forbidden", 403)
	}

	recipe.Creator = originalRecipe.Creator

	dbRecipe := recipe.toDB()
	err = api.db.UpdateRecipe(&dbRecipe)
	if err != nil {
		return err
	}
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

	if !userCanModifyRecipe(r.user, &recipe) {
		return r.writeError("Forbidden", 403)
	}

	err = api.db.DeleteRecipe(id)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) registerUser(r request) error {
	var reg UserRegistration
	err := json.NewDecoder(r.req.Body).Decode(&reg)
	if err != nil {
		return err
	}

	if !r.validateStruct(&reg) {
		return nil
	}

	if userExists := api.db.GetUserByName(reg.Username); userExists != nil {
		return r.writeError("User with this name already exists", 400)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := db.User{
		Username:    reg.Username,
		DisplayName: reg.DisplayName,
		Password:    hash,
		IsAdmin:     false,
	}

	err = api.db.AddUser(&user)
	if err != nil {
		return err
	}
	r.code = 201
	return nil
}

func (api *API) login(r request) error {
	var req LoginRequest
	err := json.NewDecoder(r.req.Body).Decode(&req)
	if err != nil {
		return err
	}

	user := api.db.GetUserByName(req.Username)
	if user == nil {
		return r.writeError("Login failed", 401)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return r.writeError("Login failed", 401)
	}

	sessionID, err := api.db.NewSession(user.ID)
	if err != nil {
		return err
	}

	r.code = 201
	return r.writeJson(map[string]string{
		"SessionID": sessionID,
	})
}

func (api *API) logout(r request) error {
	if err := api.db.DeleteSession(r.token); err != nil {
		return err
	}
	return nil
}

func (api *API) getLoggedInUser(r request) error {
	if r.user == nil {
		return r.writeJson(nil)
	}
	return r.writeJson(User{
		ID:          r.user.ID,
		DisplayName: r.user.DisplayName,
		Username:    r.user.Username,
		IsAdmin:     r.user.IsAdmin,
	})
}
