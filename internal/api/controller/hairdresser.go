package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	hairdresser_repository "api-planning/repository"
)


func FetchHairDresser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        hairdressers, err := hairdresser_repository.GetHairDresser(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des salon: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(hairdressers)
    }
}