package crypto

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/edznux/wonderxss/api"

	log "github.com/sirupsen/logrus"
)

func GetJWTToken(user api.User, key string) (string, error) {
	signingKey := []byte(key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,       // Unique
		"user_name": user.Username, // This help for UI, I don't want to make http req to get the username...
		"role":      "admin",       // Maybe we will add authz later on.
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func VerifyJWTToken(tokenString string, key string) (jwt.Claims, error) {
	signingKey := []byte(key)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	data := token.Claims.(jwt.MapClaims)
	ts := int64(data["exp"].(float64))
	now := time.Now().Unix()
	diff := (ts - now)
	log.Debug("Token exp:", ts, "now is: ", now, "diff :", diff)

	isValid := data.VerifyExpiresAt(now, true)
	if !isValid {
		return nil, fmt.Errorf("Token is expired")
	}
	return token.Claims, err
}
