package db

import (
	"time"

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

type RecipeInfo struct {
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Name             string
	ShortDescription string
	Image            *Image
	ImageID          uint
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

type User struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Username    string
	DisplayName string
	IsAdmin     bool
	Password    []byte
}

type Session struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	UserID    uint
}

type Image struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
	UserID    uint
	Size      uint
}
