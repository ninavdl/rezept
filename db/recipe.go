package db

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type Ingredient struct {
	ID       uint
	RecipeID int
	Name     string
	Amount   decimal.Decimal `gorm:"type:VARCHAR(8)"`
	Unit     string
	Note     string
}

type Step struct {
	ID       uint
	RecipeID uint
	Text     string
	Image    *Image
	ImageID  uint
}

type Recipe struct {
	RecipeInfo
	Servings    uint
	Description string
	Creator     *User
	CreatorID   uint
	Ingredients []Ingredient
	Steps       []Step
	Tags        []Tag
}

type Tag struct {
	ID       uint
	RecipeID uint
	Tag      string
}

func (db *DB) GetRecipe(id uint) *Recipe {
	var recipe Recipe
	if db.db.Where("id = ?", id).Preload("Image").Preload("Ingredients").Preload("Steps").Preload("Steps.Image").Preload("Creator").Preload("Tags").First(&recipe).RecordNotFound() {
		return nil
	}
	return &recipe
}

func (db *DB) PutRecipe(recipe *Recipe) error {
	if recipe.Creator != nil {
		// do not overwrite user in database
		recipe.CreatorID = recipe.Creator.ID
		recipe.Creator = nil
	}

	return db.db.Create(recipe).Error
}

func (db *DB) UpdateRecipe(recipe *Recipe) error {
	if recipe.ID == 0 {
		return errors.New("recipe must contain valid ID")
	}

	if recipe.Creator != nil {
		// do not overwrite user in database
		recipe.CreatorID = recipe.Creator.ID
		recipe.Creator = nil
	}

	return db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("recipe_id = ?", recipe.ID).Delete(Step{}).Error; err != nil {
			return err
		}

		if err := tx.Where("recipe_id = ?", recipe.ID).Delete(Ingredient{}).Error; err != nil {
			return err
		}

		if err := tx.Where("recipe_id = ?", recipe.ID).Delete(Tag{}).Error; err != nil {
			return err
		}

		if err := tx.Save(recipe).Error; err != nil {
			return err
		}

		return nil
	})
}

func (db *DB) DeleteRecipe(id int) error {
	return db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("recipe_id = ?", id).Delete(Step{}).Error; err != nil {
			return err
		}

		if err := tx.Where("recipe_id = ?", id).Delete(Ingredient{}).Error; err != nil {
			return err
		}

		if err := tx.Where("recipe_id = ?", id).Delete(Tag{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", id).Delete(Recipe{}).Error; err != nil {
			return err
		}

		return nil
	})
}
