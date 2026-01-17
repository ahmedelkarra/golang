package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lpernett/godotenv"
)

func JwtGen(id string) (string, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

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
