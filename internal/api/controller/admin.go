package controller


import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	admin_repository "api-planning/repository"
)


func FetchAdmin(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        admins, err := admin_repository.GetAdmin(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des admin: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(admins)
    }
}