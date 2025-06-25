package util

import (
	"app/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecretKey = []byte(config.Config("ACCESS_JWT_SECRET"))
var refreshSecretKey = []byte(config.Config("REFRESH_JWT_SECRET"))

func GenerateAccessJWT(userID uint, username string, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       userID,
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString(accessSecretKey)
}

func GenerateRefreshJWT(userID uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       userID,
		"username": username,
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

	if !ok {
		return nil, errors.New("access token 解析失敗")
	}

	return claims, err
}

func ParseRefreshJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshSecretKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("refresh token 解析失敗")
	}

	userID := uint(claims["ID"].(float64))

	return userID, err
}
