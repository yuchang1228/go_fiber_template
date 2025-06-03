package util

import (
	"app/config"
	"app/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecretKey = []byte(config.Config("ACCESS_JWT_SECRET"))
var refreshSecretKey = []byte(config.Config("REFRESH_JWT_SECRET"))

func GenerateAccessJWT(user *model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString(accessSecretKey)
}

func GenerateRefreshJWT(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(refreshSecretKey)
}

func ParseAccessJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return accessSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}

	return claims, err
}

func ParseRefreshJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}

	return claims, err
}
