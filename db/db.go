package db

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) (*DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	db.AutoMigrate(&Recipe{}, &Ingredient{}, &Step{}, &User{}, &Session{}, &Tag{}, &Image{})

	return &DB{
		db: db,
	}, nil

}

type DB struct {
	db *gorm.DB
}

func (db *DB) Close() {
	db.db.Close()
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
	db.db.Scopes(filter.apply).Table("recipes").Count(&count)
	return count
}

func (db *DB) GetRecipes(filter RecipeFilter, offset, count uint) ([]RecipeInfo, error) {
	var recipeInfos []RecipeInfo
	err := db.db.Scopes(filter.apply).Table("recipes").Preload("Image").Limit(count).Offset(offset).Find(&recipeInfos).Error
	if err != nil {
		return nil, err
	}

	return recipeInfos, nil
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

func (db *DB) GetUserById(id uint) *User {
	var user User
	if db.db.Where("id = ?", id).First(&user).RecordNotFound() {
		return nil
	}

	return &user
}

func (db *DB) GetUserByName(name string) *User {
	var user User
	if db.db.Where("lower(username) = lower(?)", name).First(&user).RecordNotFound() {
		return nil
	}

	return &user
}

func (db *DB) AddUser(user *User) error {
	return db.db.Create(user).Error
}

func (db *DB) NewSession(userid uint) (string, error) {
	id, err := GenerateRandomString(64)
	if err != nil {
		return "", err
	}
	err = db.db.Create(&Session{
		ID:     id,
		UserID: userid,
	}).Error
	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *DB) GetUserBySession(sessionID string) *User {
	var sess Session
	if db.db.Where("id = ?", sessionID).Preload("User").First(&sess).RecordNotFound() {
		return nil
	}

	return &sess.User
}

func (db *DB) DeleteSession(sessionID string) error {
	return db.db.Where("id = ?", sessionID).Delete(&Session{}).Error
}

func (db *DB) GetImageByID(id uint) *Image {
	var img Image
	if db.db.Where("id = ?", id).First(&img).RecordNotFound() {
		return nil
	}

	return &img
}

func (db *DB) PutImage(image *Image) error {
	if image.User != nil {
		// do not overwrite user in database
		image.UserID = image.User.ID
		image.User = nil
	}

	return db.db.Create(image).Error
}

func (db *DB) DeleteImage(id uint) error {
	return db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(Step{}).Where("image_id = ?", id).Update("image_id", 0).Error; err != nil {
			return err
		}

		if err := tx.Model(Recipe{}).Where("image_id = ?", id).Update("image_id", 0).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", id).Delete(Image{}).Error; err != nil {
			return err
		}

		return nil
	})
}
