package db

import "time"

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

func (db *DB) GetUserCount() int {
	var count int
	db.db.Table("users").Count(&count)
	return count
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
