package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Image struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
	UserID    uint
	Size      uint
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

func (db *DB) ListAndDeleteUnusedImages() ([]uint, error) {
	var ids []uint
	err := db.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("images").
			Where("id NOT IN (?)", tx.Table("steps").Select("image_id").QueryExpr()).
			Where("id NOT IN (?)", tx.Table("recipes").Select("image_id").QueryExpr()).
			Pluck("id", &ids).Error
		if err != nil {
			return err
		}

		fmt.Println(ids)

		err = tx.Where("id IN (?)", ids).Delete(Image{}).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ids, nil
}
