package crypto

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage/models"
)

var cfg config.Config

func init() {
	var err error
	cfg, err = config.Load("")
	if err != nil {
		log.Fatal(err)
	}
}

func GetJWTToken(user models.User) (string, error) {
	signingKey := cfg.JWTToken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    "admin", // maybe we will add authz later on.
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func VerifyJWTToken(tokenString string) (jwt.Claims, error) {
	signingKey := cfg.JWTToken
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
