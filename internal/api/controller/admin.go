package controller

import (
	"api-planning/model"
	admin_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func FetchAdminById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			log.Printf("ID manquant dans l'URL")
			http.Error(w, "ID manquant", http.StatusBadRequest)
			return
		}

		adminID, err := admin_repository.GetAdminById(db, id)
		if err != nil {
			log.Printf("Erreur lors de la récupération du coiffeur: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}


		json.NewEncoder(w).Encode(adminID)
	}
}
	

func CreateAdminHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var admin model.Admin

		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		adminID, err := admin_repository.CreateAdmin(db, admin)
		if err != nil {
			log.Printf("Erreur lors de la création de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(adminID)
	}
}

func UpdateAdminHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var admin model.Admin
		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		adminID, err := admin_repository.UpdateAdmin(db, admin)
		if err != nil {
			log.Printf("Erreur lors de la mise à jour de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(adminID)
	}
}


func DeleteAdminHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")
		if id == "" {
			log.Printf("ID manquant dans l'URL")
			http.Error(w, "ID manquant", http.StatusBadRequest)
			return
		}

		_, err := admin_repository.DeleteAdmin(db, id)
		if err != nil {
			log.Printf("Erreur lors de la récupération du coiffeur: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Coiffeur supprimé avec succès"))
	}
}