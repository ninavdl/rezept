package api

import (
	"encoding/json"

	"github.com/sour-dough/rezept-backend/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint
	Username    string
	DisplayName string
	IsAdmin     bool
}

func (u *User) toDB() db.User {
	return db.User{
		ID:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		IsAdmin:     u.IsAdmin,
	}
}

func newUser(u *db.User) User {
	return User{
		ID:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		IsAdmin:     u.IsAdmin,
	}
}

type UserRegistration struct {
	Username    string `validate:"required"`
	DisplayName string
	Password    string `validate:"required,min=8"`
}

type LoginRequest struct {
	Username string
	Password string
}

func (user *User) CanModifyRecipe(recipe *Recipe) bool {
	if user.IsAdmin {
		return true
	}

	return recipe != nil && recipe.Creator != nil && recipe.Creator.ID == user.ID
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
