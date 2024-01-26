package controller

import (
	"api-planning/internal/utils"
	"api-planning/model"
	customer_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FetchCustomer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := customer_repository.GetCustomer(db)
		if err != nil {
			utils.HandleError(w,"Erreur lors de la récupération des clients: %v", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(customers)
	}
}

func FetchCustomerById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			log.Printf("ID manquant dans l'URL")
			http.Error(w, "ID manquant", http.StatusBadRequest)
			return
		}

		customer, err := customer_repository.GetCustomerByID(db, id)
		if err != nil {
			log.Printf("Erreur lors de la récupération du client: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(customer)
	}
}

func CreateCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer model.Customer

		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		customerID, err := customer_repository.CreateCustomer(db, customer)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customerID)
	}
}

func UpdateCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "ID manquant dans l'URL", http.StatusBadRequest)
			return
		}

		var customer model.Customer

		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			log.Printf("Erreur lors de la récupération ddu client : %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		customer.ID = id

		updatedCustomer, err := customer_repository.UpdateCustomer(db, customer)
		if err != nil {
			log.Printf("Erreur lors de la création du client: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(updatedCustomer)
	}
}

func DeleteCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			http.Error(w, "ID manquant dans la requête", http.StatusBadRequest)
			return
		}

		_, err := customer_repository.DeleteCustomer(db, id)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Client supprimé avec succès"))
	}
}
