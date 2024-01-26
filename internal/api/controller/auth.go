package controller

import (
	"api-planning/internal/utils"
	"api-planning/model"
	customer_repository "api-planning/repository"

	"database/sql"
	"encoding/json"
	"net/http"
)



func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
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
		

		
		token, err := utils.GenerateUserAccessToken(customerID)
		
		if err != nil{
			http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		}

		response := map[string]interface{}{
			"customerID": customerID,
			"token":      token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		

	}
}
