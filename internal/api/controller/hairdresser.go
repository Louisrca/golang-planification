package controller

import (
	"api-planning/model"
	hairdresser_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func FetchHairDresser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hairdressers, err := hairdresser_repository.GetHairDresser(db)
		if err != nil {
			log.Printf("Erreur lors de la récupération des coiffeurs: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hairdressers)
	}
}

func FetchHairDresserById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			log.Printf("ID manquant dans l'URL")
			http.Error(w, "ID manquant", http.StatusBadRequest)
			return
		}

		hairdresser, err := hairdresser_repository.GetHairDresserByID(db, id)
		if err != nil {
			log.Printf("Erreur lors de la récupération du coiffeur: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hairdresser)
	}
}

func CreateHairDresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hairdresser model.Hairdresser

		err := json.NewDecoder(r.Body).Decode(&hairdresser)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		hairdresserID, err := hairdresser_repository.CreateHairDresser(db, hairdresser)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hairdresserID)
	}
}

func UpdateHairDresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "ID manquant dans l'URL", http.StatusBadRequest)
			return
		}

		var hairdresser model.Hairdresser

		err := json.NewDecoder(r.Body).Decode(&hairdresser)
		if err != nil {
			log.Printf("Erreur lors de la récupération du coiffeur: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		hairdresser.ID = id

		updatedHairdresser, err := hairdresser_repository.UpdateHairDresser(db, hairdresser)
		if err != nil {
			log.Printf("Erreur lors de la création du coiffeur: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(updatedHairdresser)
	}
}

func DeleteHairDresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			http.Error(w, "ID manquant dans la requête", http.StatusBadRequest)
			return
		}

		_, err := hairdresser_repository.DeleteHairDresser(db, id)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Coiffeur supprimé avec succès"))
	}
}
