package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)




func ValidateTokenAndGetID(tokenString string) (string,string, bool) {
    jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(jwtAccessSecret), nil
    })

    if err != nil {
        fmt.Println("err", err)
        return "", "", false
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        
        userID := claims["userID"].(string) 
		role := claims["role"].(string)
        return userID,role, true
    }

    return "","", false
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := ""
        authHeader := r.Header.Get("Authorization")
        if authHeader != "" {
            parts := strings.Split(authHeader, "Bearer ")
            if len(parts) == 2 {
                tokenString = parts[1]
            }
        }

        userID, _, valid := ValidateTokenAndGetID(tokenString)
        if tokenString == "" || !valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), "userID", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

