package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	slot_repository "api-planning/repository"
)


func FetchSlot(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        slots, err := slot_repository.GetSlot(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des salon: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(slots)
    }
}