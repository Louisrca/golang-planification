package controller

import (
	"api-planning/internal/utils"
	"api-planning/model"
	admin_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)



func FetchAdmin(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        admins, err := admin_repository.GetAdmin(db)
         if err != nil {
			utils.HandleError(w , "Erreur lors de la récupération des admin: %v", err, http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(admins)
    }
}

func FetchAdminById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		// err := utils.FindByID(db, &model.Admin{}, "admin", id)
	
		if id == "" {
			utils.HandleError(w,"ID manquant dans l'URL ou ID incorrect", nil, http.StatusInternalServerError )
			return
		}

		adminID, err := admin_repository.GetAdminById(db, id)
		if err != nil {
			utils.HandleError(w,"Erreur lors de la récupération du admin: %v", err, http.StatusInternalServerError )
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
			utils.HandleError(w,"Requête invalide", err, http.StatusInternalServerError)
			return
		}

		adminID, err := admin_repository.CreateAdmin(db, admin)
		if err != nil {
			utils.HandleError(w,"Erreur lors de la création de l'admin: %v", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(adminID)
	}
}

func UpdateAdminHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var admin model.Admin
		id := chi.URLParam(r, "id")

		err := utils.FindByID(db, admin, "admin", id)
		
		if id == "" || err == nil  {
			utils.HandleError(w,"ID manquant dans l'URL", nil, http.StatusInternalServerError)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&admin)

		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération de l'admin", err, http.StatusInternalServerError)
			return
		}

		adminID, err := admin_repository.UpdateAdmin(db, admin)

		if err != nil {
			utils.HandleError(w,"Erreur lors de la mise à jour de l'admin", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(adminID)
	}
}


func DeleteAdminHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")
		if id == "" {
			utils.HandleError(w,"ID manquant dans l'URL", nil, http.StatusInternalServerError)
			return
		}

		_, err := admin_repository.DeleteAdmin(db, id)
		if err != nil {
			utils.HandleError(w,"Erreur lors de la suppression de l'admin", err, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Admin supprimé avec succès"))
	}
}