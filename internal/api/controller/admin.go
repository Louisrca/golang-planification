package controller

import (
	"api-planning/model"
	admin_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		id := r.URL.Query().Get("id")
		adminID, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		admin, err := admin_repository.GetAdminById(db, adminID)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(admin)
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

		id := r.URL.Query().Get("id")
		adminID, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		admin, err := admin_repository.DeleteAdmin(db, adminID)
		if err != nil {
			log.Printf("Erreur lors de la suppression de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(admin)
	}
}