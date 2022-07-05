package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateTokenFromUser(userId uint) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	exp := time.Now().Add(time.Hour * 24 * 30).Unix()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["userId"] = userId
	claims["exp"] = exp

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println(err)
		return tokenString, err
	}

	return tokenString, nil
}

func DecodeTokenFromUser(tokenString string) (jwt.MapClaims, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !token.Valid {
		return nil, err
	}

	return claims, err
}
