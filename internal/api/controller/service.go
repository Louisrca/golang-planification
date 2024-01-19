package controller

import (
	"api-planning/model"
	service_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func FetchService(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services, err := service_repository.GetService(db)
		if err != nil {
			log.Printf("Erreur lors de la récupération des salon: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(services)
	}
}

func FetchServiceById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		service_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de service: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		service, err := service_repository.GetServiceByID(db, service_id)
		if err != nil {
			log.Printf("Erreur lors de la récupération d: %v", err)
			http.Error(w, http.StatusText(500), 500)
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
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		_, err = service_repository.CreateService(db, service)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(service)
	}
}
func UpdateServiceHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var service model.Service

		err := json.NewDecoder(r.Body).Decode(&service)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		_, err = service_repository.UpdateService(db, service)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(service)
	}
}

func DeleteServiceHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		serviceID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "ID de service invalide", http.StatusBadRequest)
			return
		}

		_, err = service_repository.DeleteService(db, serviceID)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service supprimé avec succès"))
	}
}
