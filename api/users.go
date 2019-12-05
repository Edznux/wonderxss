package api

import (
	"errors"
	"fmt"

	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
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

func GetUserByName(name string) (models.User, error) {
	user, err := store.GetUserByName(name)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(username, password string) (models.User, error) {
	u := models.User{}
	//We don't want empty username
	if username == "" {
		return u, errors.New("Invalid username")
	}

	// Simplest password policy
	if len(password) < 10 {
		return u, errors.New("Invalid password")
	}

	existingUser, err := GetUserByName(username)
	if err != nil {
		// If the error is just an empty response, ignore
		if err != models.NoSuchItem {
			fmt.Println(err)
			return u, errors.New("Database error")
		}
	}

	// Yes, I know, user enum', don't care / will fix by other means (rate limit, captcha...)
	if existingUser.ID != "" || existingUser.Username != "" {
		return u, errors.New("user already exist")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}

	u.ID = uuid.New().String()
	u.Username = username
	u.Password = string(hashedPass)

	user, err := store.CreateUser(u)
	if err != nil {
		return u, err
	}
	return user, nil
}
