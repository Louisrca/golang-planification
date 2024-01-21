package controller

import (
	"api-planning/internal/utils"
	"api-planning/model"
	booking_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func FetchBooking(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        bookings, err := booking_repository.GetBooking(db)
         if err != nil {
			utils.HandleError(w,"Erreur lors de la récupération des réservations: %v", nil, http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(bookings)
    }
}

func FetchBookingById(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id ==""{
			utils.HandleError(w,"ID manquant dans l'URL", nil, http.StatusInternalServerError)
            return
		}
		
        bookings, err := booking_repository.GetBookingById(db, id)
         if err != nil {
			utils.HandleError(w,"Erreur lors de la récupération des réservations: %v", err, http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(bookings)
    }
}


func CreateBookingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Création d'une instance pour stocker les données décodées
		var booking model.Booking

		// Décodage du corps de la requête
		err := json.NewDecoder(r.Body).Decode(&booking)
		if err != nil {
			utils.HandleError(w,"Requête invalide", err, http.StatusInternalServerError)
			return
		}

		// Appel de la fonction pour créer une nouvelle réservation
		bookingID, err := booking_repository.CreateBooking(db, booking)
		if err != nil {
			utils.HandleError(w,"Erreur lors de la création de la réservation: %v", err, http.StatusInternalServerError)
			return
		}

		// Envoie d'une réponse réussie
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bookingID)
	}

}


func UpdateBookingHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Création d'une instance pour stocker les données décodées
        var updatedBooking model.Booking

        // Décodage du corps de la requête
        err := json.NewDecoder(r.Body).Decode(&updatedBooking)
        if err != nil {
			utils.HandleError(w,"Requête invalide", err, http.StatusInternalServerError)
            return
        }

        // Appel de la fonction pour mettre à jour la réservation
        bookingID, err := booking_repository.UpdateBooking(db, updatedBooking)
        if err != nil {
			utils.HandleError(w,"Erreur lors de la mise à jour de la réservation: %v", err, http.StatusInternalServerError)
            return
        }

        // Envoie d'une réponse réussie
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(bookingID)
    }
}



func DeleteBookingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupération de l'id de la réservation à supprimer
		id := chi.URLParam(r, "id")
		if id == "" {
			utils.HandleError(w,"Requête invalide", nil, http.StatusInternalServerError)
			return
		}

		// Appel de la fonction pour supprimer la réservation
		err := booking_repository.DeleteBooking(db, id)
		if err != nil {
			utils.HandleError(w,"Erreur lors de la suppression de la réservation: %v", err, http.StatusInternalServerError)
			return
		}

		// Envoie d'une réponse réussie
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Réservation supprimée avec succès"))
	}
}