package db

import (
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
