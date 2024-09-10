package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key string = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(userID int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userID,
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(key))
}
