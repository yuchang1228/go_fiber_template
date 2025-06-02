package util

import (
	"app/config"
	"app/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecretKey = []byte(config.Config("ACCESS_JWT_SECRET"))
var refreshSecretKey = []byte(config.Config("REFRESH_JWT_SECRET"))

func GenerateAccessJWT(user *model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString(accessSecretKey)
}

func GenerateRefreshJWT(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(refreshSecretKey)
}
