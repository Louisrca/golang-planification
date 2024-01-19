package controller

import (
	"api-planning/model"
	booking_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)


func FetchBooking(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        bookings, err := booking_repository.GetBooking(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des clients: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(bookings)
    }
}


func CreateBookingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Création d'une instance pour stocker les données décodées
		var newBooking model.Booking

		// Décodage du corps de la requête
		err := json.NewDecoder(r.Body).Decode(&newBooking)
		if err != nil {
			log.Printf("Erreur lors de la lecture de la requête: %v", err)
			http.Error(w, http.StatusText(400), 400) // Bad Request
			return
		}

		// Appel de la fonction pour créer une nouvelle réservation
		err = booking_repository.CreateBooking(db, newBooking)
		if err != nil {
			log.Printf("Erreur lors de la création de la réservation: %v", err)
			http.Error(w, http.StatusText(500), 500) // Internal Server Error
			return
		}

		// Envoie d'une réponse réussie
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBooking)
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
        err = booking_repository.UpdateBooking(db, updatedBooking)
        if err != nil {
            log.Printf("Erreur lors de la mise à jour de la réservation: %v", err)
            http.Error(w, http.StatusText(500), 500) // Internal Server Error
            return
        }

        // Envoie d'une réponse réussie
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(updatedBooking)
    }
}



func DeleteBookingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupération de l'id de la réservation à supprimer
		bookingID := r.URL.Query().Get("id")
		if bookingID == "" {
			log.Printf("Erreur lors de la lecture de la requête: %v", bookingID)
			http.Error(w, http.StatusText(400), 400) // Bad Request
			return
		}

		// Appel de la fonction pour supprimer la réservation
		err := booking_repository.DeleteBooking(db, bookingID)
		if err != nil {
			log.Printf("Erreur lors de la suppression de la réservation: %v", err)
			http.Error(w, http.StatusText(500), 500) // Internal Server Error
			return
		}

		// Envoie d'une réponse réussie
		w.WriteHeader(http.StatusOK)
	}
}