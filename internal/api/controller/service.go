package controller

import (
	"api-planning/model"
	service_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"fmt"
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

        // Décoder le corps de la requête en un objet service
        err := json.NewDecoder(r.Body).Decode(&service)
        if err != nil {
            http.Error(w, "Requête invalide", http.StatusBadRequest)
            return
        }

        // Mise à jour du service et récupération des informations mises à jour
        updatedService, err := service_repository.UpdateService(db, service)
        if err != nil {
            http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
            return
        }

        // Envoyer les informations mises à jour en réponse
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(updatedService)
    }
}


func DeleteServiceHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")
		fmt.Println(id)
        if id == "" {
            http.Error(w, "ID manquant dans la requête", http.StatusBadRequest)
            return
        }

        _, err := service_repository.DeleteService(db, id)
        if err != nil {
            http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Service supprimé avec succès"))
    }
}



