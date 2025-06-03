package helper

import (
	"e-novel/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 100000000).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
