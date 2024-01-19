package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	hair_salon_repository "api-planning/repository"
)


func FetchHairSalon(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        hair_salons, err := hair_salon_repository.GetHairSalon(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des salon: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(hair_salons)
    }
}