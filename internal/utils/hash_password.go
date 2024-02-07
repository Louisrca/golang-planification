package utils

import (
	"crypto/sha256"
	"encoding/base64"


)

func HashPassword(password string) string {

	hash := sha256.New()

	hash.Write([]byte(password))

	hashedPassword := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	return hashedPassword
}


func CheckPasswordHash(password, hashedPassword string) bool {
    return HashPassword(password) == hashedPassword
}