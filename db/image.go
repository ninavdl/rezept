package db

import (
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
