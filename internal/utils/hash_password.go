package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashPassword(password string) string {
	// Créez un nouveau hachage SHA-256
	hash := sha256.New()

	// Écrivez le mot de passe dans le hachage
	hash.Write([]byte(password))

	// Récupérez la somme de contrôle finale et encodez-la en base64
	hashedPassword := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	return hashedPassword
}