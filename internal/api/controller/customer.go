package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	customer_repository "api-planning/repository"
)


func FetchCustomer(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        customers, err := customer_repository.GetCustomer(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des clients: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(customers)
    }
}