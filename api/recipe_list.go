package api

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/sour-dough/rezept-backend/db"
)

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
