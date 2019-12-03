package api

import (
	"errors"
	"github.com/edznux/wonderxss/storage/models"
	"golang.org/x/crypto/bcrypt"
)

func VerifyUserPassword(loginParam, passwordParam string) (models.User, error) {
	var err error
	var user models.User

	// Check for empty user / password
	if loginParam == "" || passwordParam == "" {
		return user, errors.New("Empty user or password")
	}
	// Get user model by its name
	user, err = store.GetUserByName(loginParam)
	if err != nil {
		return user, errors.New("Invalid user or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordParam))
	if err != nil {
		return user, errors.New("Invalid user or password")
	}

	return user, nil
}

func CreateUser(username, password string) (models.User, error){
	user := models.User{}
	return user, nil
}