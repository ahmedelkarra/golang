package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JwtGen(id string) (string, error) {
	secretKey := []byte(os.Getenv("secretKey"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour + 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
