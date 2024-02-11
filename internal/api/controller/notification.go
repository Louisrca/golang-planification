package controller

import (
	admin_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func FetchNotificationByAdminId(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			log.Printf("ID manquant dans l'URL")
			http.Error(w, "ID manquant", http.StatusBadRequest)
			return
		}

		notifications, err := admin_repository.GetNotificationByAdminID(db, id)
		if err != nil {
			log.Printf("Erreur lors de la récupération des notidications de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(notifications)
	}
}
