package controller

import (
	"api-planning/internal/utils"
	"api-planning/model"
	slot_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FetchSlot(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slots, err := slot_repository.GetSlot(db)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération des créneaux", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(slots)
	}
}

func FetchSlotById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			utils.HandleError(w, "ID de créneau manquant", nil, http.StatusBadRequest)
			return
		}

		slot, err := slot_repository.GetSlotByID(db, id)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération du créneau", err, http.StatusInternalServerError)
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
			utils.HandleError(w, "Requête invalide", err, http.StatusBadRequest)
			return
		}

		slotId, err := slot_repository.CreateSlot(db, slot)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la création du créneau", err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(slotId)
	}
}

func UpdateSlotHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "ID manquant dans l'URL", http.StatusBadRequest)
			return
		}
		var slot model.Slot

		err := json.NewDecoder(r.Body).Decode(&slot)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupartion du créneau", err, http.StatusInternalServerError)
			return
		}
		slot.ID = id

		updateSlot, err := slot_repository.UpdateSlot(db, slot)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la mise à jour du créneau", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updateSlot)
	}
}

func DeleteSlotHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			utils.HandleError(w, "ID de créneau manquant", nil, http.StatusBadRequest)
			return
		}

		_, err := slot_repository.DeleteSlot(db, id)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la suppression du créneau", err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Créneau supprimé avec succès"))
	}
}
