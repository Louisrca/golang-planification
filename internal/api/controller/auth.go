package controller

import (
	// "api-planning/internal/middleware"
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

		token, err := utils.GenerateUserAccessToken(customerID.ID, "customer")

		if err != nil {
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
	return func(w http.ResponseWriter, r *http.Request) {
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

		token, err := utils.GenerateUserAccessToken(adminID.ID, "admin")

		if err != nil {
			http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		}

		response := map[string]interface{}{
			"adminID": adminID,
			"token":   token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}
func RegisterHaidresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		token, err := utils.GenerateUserAccessToken(hairdresserID.ID, "hairdresser")

		if err != nil {
			http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		}

		response := map[string]interface{}{
			"hairdresserID": hairdresserID,
			"token":         token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func LoginCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds model.UserCredentials

		// Décoder le JSON du corps de la requête
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		customer, err := customer_repository.GetCustomerByEmail(db, creds.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				// Utilisateur non trouvé
				http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
			} else {
				// Autre erreur de base de données
				http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			}
			return
		}

		if !utils.CheckPasswordHash(creds.Password, customer.Password) {
			http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
			return
		}

		tokenString, err := utils.GenerateUserAccessToken(customer.Email, "customer")
		if err != nil {
			utils.HandleError(w, "Erreur lors de la génération du JWT", err, http.StatusInternalServerError)
			return
		}

		// Renvoyer le token JWT
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
func LoginAdminHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds model.UserCredentials

		// Décoder le JSON du corps de la requête
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		admin, err := admin_repository.GetAdminByEmail(db, creds.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				// Utilisateur non trouvé
				http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
			} else {
				// Autre erreur de base de données
				http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			}
			return
		}

		if !utils.CheckPasswordHash(creds.Password, admin.Password) {
			http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
			return
		}

		tokenString, err := utils.GenerateUserAccessToken(admin.Email, "admin")
		if err != nil {
			utils.HandleError(w, "Erreur lors de la génération du JWT", err, http.StatusInternalServerError)
			return
		}

		// Renvoyer le token JWT
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

func LoginHairdresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds model.UserCredentials

		// Décoder le JSON du corps de la requête
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hairdresser, err := hairdresser_repository.GetHairDresserByEmail(db, creds.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				// Utilisateur non trouvé
				http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
			} else {
				// Autre erreur de base de données
				http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			}
			return
		}

		if !utils.CheckPasswordHash(creds.Password, hairdresser.Password) {
			http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
			return
		}

		tokenString, err := utils.GenerateUserAccessToken(hairdresser.Email, "hairdresser")
		if err != nil {
			utils.HandleError(w, "Erreur lors de la génération du JWT", err, http.StatusInternalServerError)
			return
		}

		// Renvoyer le token JWT
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
