package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	booking_repository "api-planning/repository"
)


func FetchBooking(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        bookings, err := booking_repository.GetBooking(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des clients: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(bookings)
    }
}