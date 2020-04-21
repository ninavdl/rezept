package db

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type RecipeInfo struct {
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Name             string
	ShortDescription string
	Image            *Image
	ImageID          uint
	Published        bool `gorm:"default:true"`
}

type RecipeFilter struct {
	Tags     []string
	Keywords []string
	User     uint
}

func (f RecipeFilter) apply(db *gorm.DB) *gorm.DB {
	if f.Tags != nil {
		db = db.Where("id IN (?)", db.Table("tags").Select("recipe_id").Where("tag IN (?)", f.Tags).QueryExpr())
	}
	if f.Keywords != nil {
		q := make([]string, len(f.Keywords))
		params := make([]interface{}, len(f.Keywords)*3)
		for i, keyword := range f.Keywords {
			q[i] = "name LIKE ? OR description LIKE ? OR short_description LIKE ?"
			params[3*i] = "%" + keyword + "%"
			params[3*i+1] = "%" + keyword + "%"
			params[3*i+2] = "%" + keyword + "%"
		}
		queryString := strings.Join(q, " OR ")
		db = db.Where(queryString, params...)
	}
	if f.User != 0 {
		db = db.Where("creator_id = ?", f.User)
	}

	return db
}

func (db *DB) CountRecipes(filter RecipeFilter) uint {
	var count uint
	db.db.Scopes(filter.apply).Where("published = 1").Table("recipes").Count(&count)
	return count
}

func (db *DB) GetRecipes(filter RecipeFilter, offset, count uint) ([]RecipeInfo, error) {
	var recipeInfos []RecipeInfo
	err := db.db.Scopes(filter.apply).Where("published = 1").Table("recipes").Preload("Image").Limit(count).Offset(offset).Find(&recipeInfos).Error
	if err != nil {
		return nil, err
	}

	return recipeInfos, nil
}

func (db *DB) GetDrafts(user uint) ([]RecipeInfo, error) {
	var recipeInfos []RecipeInfo
	err := db.db.Where("published = 0").Where("creator_id = ?", user).Table("recipes").Preload("Image").Find(&recipeInfos).Error
	if err != nil {
		return nil, err
	}

	return recipeInfos, nil
}
