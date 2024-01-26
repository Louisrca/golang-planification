package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateUserAccessToken(model interface{}) (string, error) {
	// Load .env file

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	jwtSigningKey := os.Getenv("JWT_SIGNING_KEY")

	jwtCreated := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerID": model,
	})

	token, err := jwtCreated.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	fmt.Println(token)

	return token, nil
}
