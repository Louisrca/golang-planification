package controller

import (
	"api-planning/model"
	customer_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func FetchCustomer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := customer_repository.GetCustomer(db)
		if err != nil {
			log.Printf("Erreur lors de la récupération des clients: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(customers)
	}
}

func FetchCustomerById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		customerID, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		customer, err := customer_repository.GetCustomerByID(db, customerID)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(customer)
	}
}

func CreateCustomerHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var customer model.Customer

        // Décoder le corps de la requête en un objet customer
        err := json.NewDecoder(r.Body).Decode(&customer)
        if err != nil {
            http.Error(w, "Requête invalide", http.StatusBadRequest)
            return
        }

        // Créer le client dans la base de données
        customerID, err := customer_repository.CreateCustomer(db, customer)
        if err != nil {
            http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
            return
        }

        // Envoyer l'ID du client créé en réponse
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]int64{"id": customerID})
    }
}


func UpdateCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var customer model.Customer

		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		customer_id, err := customer_repository.UpdateCustomer(db, customer)
		if err != nil {
			log.Printf("Erreur lors de la création de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(customer_id)
	}
}


func DeleteCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		customerID, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		err = customer_repository.DeleteCustomer(db, customerID)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}