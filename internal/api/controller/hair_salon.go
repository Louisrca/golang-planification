package controller

import (
	"api-planning/model"
	admin_repository "api-planning/repository"
	hair_salon_repository "api-planning/repository"
	hairdresser_repository "api-planning/repository"
	notification_respository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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

		hairdresserID, ok := r.Context().Value("userID").(string)
		if !ok || hairdresserID == "" {
			http.Error(w, "Unauthorized or missing hairdresser ID", http.StatusUnauthorized)
			return
		}

		updatedHairdresser, err := hairdresser_repository.UpdateHairdresser(db, hairdresserID, hairSalonID.ID)
		if err != nil {
			log.Printf("Erreur lors de la mise à jour du coiffeur: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Notifier l'admin
		adminID := "197f586a-6402-4590-87c3-ceacd4558b22"
		admin, err := admin_repository.GetAdminById(db, adminID)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'admin: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		notification_respository.SendNotificationToAdmin(db, admin, hairSalonID.ID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"hairSalonID":        hairSalonID,
			"updatedHairdresser": updatedHairdresser,
		})
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

func AcceptHairSalonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			http.Error(w, "ID manquant dans la requête", http.StatusBadRequest)
			return
		}

		updatedHairSalon, err := hair_salon_repository.AcceptHairSalon(db, id)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updatedHairSalon)

	}
}
