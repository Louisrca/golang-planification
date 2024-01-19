package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	category_repository "api-planning/repository"
)


func FetchCategory(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        category, err := category_repository.GetCategory(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des salon: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(category)
    }
}