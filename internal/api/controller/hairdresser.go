package controller

import (
	"api-planning/model"
	hairdresser_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FetchHairdresser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hairdressers, err := hairdresser_repository.GetHairdresser(db)
		if err != nil {
			log.Printf("Erreur lors de la récupération des coiffeurs: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hairdressers)
	}
}

func FetchHairdresserById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			log.Printf("ID manquant dans l'URL")
			http.Error(w, "ID manquant", http.StatusBadRequest)
			return
		}

		hairdresser, err := hairdresser_repository.GetHairdresserByID(db, id)
		if err != nil {
			log.Printf("Erreur lors de la récupération du coiffeur: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hairdresser)
	}
}

func CreateHairdresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hairdresser model.Hairdresser

		err := json.NewDecoder(r.Body).Decode(&hairdresser)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		hairdresserID, err := hairdresser_repository.CreateHairdresser(db, hairdresser)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hairdresserID)
	}
}


func DeleteHairdresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			http.Error(w, "ID manquant dans la requête", http.StatusBadRequest)
			return
		}

		_, err := hairdresser_repository.DeleteHairdresser(db, id)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Coiffeur supprimé avec succès"))
	}
}

