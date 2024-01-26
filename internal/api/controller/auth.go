package controller

import (
	"api-planning/internal/utils"
	"api-planning/model"
	admin_repository "api-planning/repository"
	customer_repository "api-planning/repository"
	hairdresser_repository "api-planning/repository"

	"database/sql"
	"encoding/json"
	"net/http"
)



func RegisterCustomerHandler(db *sql.DB) http.HandlerFunc {
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
		

		
		token, err := utils.GenerateUserAccessToken(customerID.ID)
		
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

func RegisterAdminHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var admin model.Admin

		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		adminID, err := admin_repository.CreateAdmin(db, admin)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}
		

		
		token, err := utils.GenerateUserAccessToken(adminID.ID)
		
		if err != nil{
			http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		}

		response := map[string]interface{}{
			"adminID": adminID,
			"token":      token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		

	}
}
func RegisterHaidresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var hairdresser model.Hairdresser

		err := json.NewDecoder(r.Body).Decode(&hairdresser)
		if err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}

		hairdresserID, err := hairdresser_repository.CreateHairDresser(db, hairdresser)
		if err != nil {
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}
		

		
		token, err := utils.GenerateUserAccessToken(hairdresserID.ID)
		
		if err != nil{
			http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		}

		response := map[string]interface{}{
			"hairdresserID": hairdresserID,
			"token":      token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		

	}
}



