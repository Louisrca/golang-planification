package controller

import (
	"api-planning/model"
	slot_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func FetchSlot(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slots, err := slot_repository.GetSlot(db)
		if err != nil {
			log.Printf("Erreur lors de la récupération des salon: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(slots)
	}
}

func FetchSlotById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		slot_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de slot: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		slot, err := slot_repository.GetSlotByID(db, slot_id)
		if err != nil {
			log.Printf("Erreur lors de la récupération d: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(slot)
	}
}

func CreateSlotHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var slot model.Slot

		err := json.NewDecoder(r.Body).Decode(&slot)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		_, err = slot_repository.CreateSlot(db, slot)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(slot)
	}
}

func UpdateSlotHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var slot model.Slot

		err := json.NewDecoder(r.Body).Decode(&slot)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		_, err = slot_repository.UpdateSlot(db, slot)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(slot)
	}
}

func DeleteSlotHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		
        if id == "" {
            http.Error(w, "ID manquant dans la requête", http.StatusBadRequest)
            return
        }

        _, err := slot_repository.DeleteSlot(db, id)
        if err != nil {
            http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
            return
        }

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Créneau supprimé avec succès"))
	}
}
