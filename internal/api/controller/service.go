package controller

import (
	"api-planning/internal/utils"
	"api-planning/model"
	service_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FetchService(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services, err := service_repository.GetService(db)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération des services", err, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(services)
	}
}

func FetchServiceById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			utils.HandleError(w, "ID de service manquant", nil, http.StatusBadRequest)
			return
		}

		service, err := service_repository.GetServiceByID(db, id)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération du service", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(service)
	}
}

func CreateServiceHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var service model.Service

		err := json.NewDecoder(r.Body).Decode(&service)
		if err != nil {
			utils.HandleError(w, "Requête invalide", err, http.StatusBadRequest)
			return
		}

		serviceID, err := service_repository.CreateService(db, service)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la création du service", err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(serviceID)
	}
}
func UpdateServiceHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			http.Error(w, "ID manquant dans l'URL", http.StatusBadRequest)
			return
		}
		var service model.Service

		err := json.NewDecoder(r.Body).Decode(&service)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération du service", err, http.StatusInternalServerError)
			return
		}
		service.ID = id

		updatedService, err := service_repository.UpdateService(db, service)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la mise à jour du service", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updatedService)
	}
}

func DeleteServiceHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Println(id)
		if id == "" {
			utils.HandleError(w, "ID de service manquant", nil, http.StatusBadRequest)
			return
		}

		_, err := service_repository.DeleteService(db, id)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la suppression du service", err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service supprimé avec succès"))
	}
}
