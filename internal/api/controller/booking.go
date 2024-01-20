package controller

import (
	"api-planning/model"
	booking_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func FetchBooking(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        bookings, err := booking_repository.GetBooking(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des réservations: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(bookings)
    }
}

func FetchBookingById(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id ==""{
			log.Printf("ID manquant dans l'URL")
            http.Error(w, "ID manquant dans l'URL", http.StatusBadRequest)
		}
		
        bookings, err := booking_repository.GetBookingById(db, id)
         if err != nil {
            log.Printf("Erreur lors de la récupération des réservations: %v", err)
            http.Error(w, http.StatusText(500), 500)
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
			log.Printf("Erreur lors de la lecture de la requête: %v", err)
			http.Error(w, http.StatusText(400), 400) // Bad Request
			return
		}

		// Appel de la fonction pour créer une nouvelle réservation
		bookingID, err := booking_repository.CreateBooking(db, booking)
		if err != nil {
			log.Printf("Erreur lors de la création de la réservation: %v", err)
			http.Error(w, http.StatusText(500), 500) // Internal Server Error
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
            log.Printf("Erreur lors de la lecture de la requête: %v", err)
            http.Error(w, http.StatusText(400), 400) // Bad Request
            return
        }

        // Appel de la fonction pour mettre à jour la réservation
        bookingID, err := booking_repository.UpdateBooking(db, updatedBooking)
        if err != nil {
            log.Printf("Erreur lors de la mise à jour de la réservation: %v", err)
            http.Error(w, http.StatusText(500), 500) // Internal Server Error
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
			log.Printf("Erreur lors de la lecture de la requête: %v", id)
			http.Error(w, http.StatusText(400), 400) // Bad Request
			return
		}

		// Appel de la fonction pour supprimer la réservation
		err := booking_repository.DeleteBooking(db, id)
		if err != nil {
			log.Printf("Erreur lors de la suppression de la réservation: %v", err)
			http.Error(w, http.StatusText(500), 500) // Internal Server Error
			return
		}

		// Envoie d'une réponse réussie
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Réservation supprimée avec succès"))
	}
}