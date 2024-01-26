package utils

import (
	"api-planning/model"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func generateUserAccessToken(customer model.Customer) (string, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	jwtSigningKey := os.Getenv("JWT_SIGNING_KEY")

	jwtCreated := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerID": customer.ID,
	})

	token, err := jwtCreated.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
