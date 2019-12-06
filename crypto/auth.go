package crypto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage/models"
)

func GetJWTToken(user models.User) (string, error) {
	signingKey := config.Current.JWTToken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    "admin", // maybe we will add authz later on.
	})
	tokenString, err := token.SignedString([]byte(signingKey))
	return tokenString, err
}

func VerifyJWTToken(tokenString string) (jwt.Claims, error) {
	signingKey := config.Current.JWTToken
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
