package local

import (
	"errors"
	"fmt"
	"log"

	"github.com/dgryski/dgoogauth"
	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/crypto"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// verifyOTP takes the user's OTP secret and the OTPToken. It will return true if it's valid.
func verifyOTP(secret string, otp string) (bool, error) {

	if len(secret) == 0 {
		return false, fmt.Errorf("Empty TOTPSecret")
	}

	otpc := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  3,
		HotpCounter: 0,
	}

	verified, err := otpc.Authenticate(otp)
	if err != nil {
		log.Println("VerifyOTP failed authenticate:", err)
		return false, err
	}

	if !verified {
		return false, fmt.Errorf("Invalid one-time password")
	}

	return true, nil
}

func (local *Local) Login(loginParam, passwordParam, otp string) (string, error) {
	var err error
	var user models.User
	var apiUser api.User

	// Check for empty user / password
	if loginParam == "" || passwordParam == "" {
		return "", errors.New("Empty user or password")
	}
	// Get user model by its name
	user, err = local.store.GetUserByName(loginParam)
	if err != nil {
		return "", errors.New("Invalid user or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordParam))
	if err != nil {
		return "", errors.New("Invalid user or password")
	}

	if user.TwoFactorEnabled {
		OTPOk, err := verifyOTP(user.TOTPSecret, otp)
		if err != nil {
			return "", err
		}
		if !OTPOk {
			return "", errors.New("Invalid OTP")
		}
	}

	jwt, err := crypto.GetJWTToken(apiUser.FromStorage(user))
	if err != nil {
		return "", err
	}
	return jwt, nil
}

// GetUserByName is a direct replica from the storage.
func (local *Local) GetUserByName(name string) (api.User, error) {
	var returnedUser api.User
	user, err := local.store.GetUserByName(name)
	if err != nil {
		return api.User{}, err
	}
	return returnedUser.FromStorage(user), nil
}

func (local *Local) GetUser(id string) (api.User, error) {
	var returnedUser api.User
	user, err := local.store.GetUser(id)
	if err != nil {
		return api.User{}, err
	}
	return returnedUser.FromStorage(user), nil
}

func (local *Local) CreateOTP(userID string, secret string, otp string) (api.User, error) {
	var returnedUser api.User
	user, err := local.store.GetUser(userID)
	if err != nil {
		return api.User{}, err
	}
	user, err = local.store.CreateOTP(user, secret)
	if err != nil {
		return api.User{}, err
	}
	// Verify the token with the secret
	isTokenVerified, err := verifyOTP(user.TOTPSecret, otp)
	if err != nil {
		return api.User{}, err
	}
	if !isTokenVerified {
		return api.User{}, err
	}
	return returnedUser.FromStorage(user), nil
}

func (local *Local) CreateUser(username, password string) (api.User, error) {
	var returnedUser api.User
	var user models.User
	//We don't want empty username
	if username == "" {
		return api.User{}, errors.New("Invalid username")
	}

	// Simplest password policy
	if len(password) < 10 {
		return api.User{}, errors.New("Invalid password")
	}

	existingUser, err := local.store.GetUserByName(username)
	if err != nil {
		// If the error is just an empty response, ignore
		if err != models.NoSuchItem {
			fmt.Println(err)
			return api.User{}, errors.New("Database error")
		}
	}

	// Yes, I know, user enum', don't care / will fix by other means (rate limit, captcha...)
	if existingUser.ID != "" || existingUser.Username != "" {
		return api.User{}, errors.New("user already exist")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return api.User{}, err
	}

	user.ID = uuid.New().String()
	user.Username = username
	user.Password = string(hashedPass)

	user, err = local.store.CreateUser(user)
	if err != nil {
		return api.User{}, err
	}
	return returnedUser.FromStorage(user), nil
}

func (local *Local) DeleteUser(id string) error {
	e := models.User{ID: id}
	err := local.store.DeleteUser(e)
	if err != nil {
		return err
	}
	return nil
}
