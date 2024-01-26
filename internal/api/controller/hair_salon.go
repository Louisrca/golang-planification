package controller

import (
	"api-planning/model"
	hair_salon_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func FetchHairSalon(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hair_salons, err := hair_salon_repository.GetHairSalon(db)
		if err != nil {
			log.Printf("Erreur lors de la récupération des salon: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hair_salons)
	}
}

func FetchHairSalonById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			log.Printf("ID manquant dans l'URL")
			http.Error(w, "ID manquant", http.StatusBadRequest)
			return
		}

		hair_salon, err := hair_salon_repository.GetHairSalonByID(db, id)
		if err != nil {
			log.Printf("Erreur lors de la récupération du salon de coiffure: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hair_salon)
	}
}

func CreateHairSalonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hairSalon model.HairSalon

		err := json.NewDecoder(r.Body).Decode(&hairSalon)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		hairSalonID, err := hair_salon_repository.CreateHairSalon(db, hairSalon)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hairSalonID)
	}
}

func UpdateHairSalonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "ID manquant dans l'URL", http.StatusBadRequest)
			return
		}

		var hairSalon model.HairSalon

		err := json.NewDecoder(r.Body).Decode(&hairSalon)
		if err != nil {
			log.Printf("Erreur lors de la récupération du salon de coiffure : %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		hairSalon.ID = id

		updatedHairSalon, err := hair_salon_repository.UpdateHairSalon(db, hairSalon)
		if err != nil {
			log.Printf("Erreur lors de la création du salon de coiffure: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(updatedHairSalon)
	}
}

func DeleteHairSalonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			http.Error(w, "ID manquant dans la requête", http.StatusBadRequest)
			return
		}

		_, err := hair_salon_repository.DeleteHairSalon(db, id)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Salon de coiffure supprimé avec succès"))
	}
}
