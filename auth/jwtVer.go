package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func JwtVer(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("secretKey"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
