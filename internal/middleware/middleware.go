package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenString string) bool {
	jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtAccessSecret), nil
	})
	fmt.Print("token", token)

	if err != nil {
		fmt.Println("err", err)
		return false
	}

	// Valider le token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Vérifier d'autres champs comme 'exp' (expiration) si nécessaire
		if exp, ok := claims["exp"].(float64); ok {
			return time.Unix(int64(exp), 0).After(time.Now())
		}
		fmt.Println("claims", claims)
		return true

	}

	return false
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrait le token de l'en-tête Authorization
		tokenString := ""
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, "Bearer ")
			if len(parts) == 2 {
				tokenString = parts[1]
			}
		}

		// Valide le token
		if tokenString == "" || !ValidateToken(tokenString) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Passe à la prochaine fonction si le token est valide
		next.ServeHTTP(w, r)
	})
}
