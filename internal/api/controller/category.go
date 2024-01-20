package controller

import (
	"api-planning/model"
	category_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func FetchCategory(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        category, err := category_repository.GetCategory(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des salon: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(category)
    }
}


func FetchCategoryById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r,"id")
		
		if id == "" {
			log.Printf("Erreur lors de la récupération de l'ID de category: %v", id)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		category, err := category_repository.GetCategoryByID(db,id)
		 if err != nil {
			log.Printf("Erreur lors de la récupération d: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(category)
	}
}


func CreateCategoryHandler(db *sql.DB) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {

		var category model.Category

		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			log.Printf("Erreur lors de la récupération de category: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		categoryID, err := category_repository.CreateCategory(db, category)
		if err != nil {
			log.Printf("Erreur lors de la création de category: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(categoryID)
	}
}

func UpdateCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Création d'une instance pour stocker les données décodées
		var category model.Category

		// Décodage du corps de la requête
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			log.Printf("Erreur lors de la lecture de la requête: %v", err)
			http.Error(w, http.StatusText(400), 400) // Bad Request
			return
		}

		// Appel de la fonction pour mettre à jour la réservation
		categoryID, err := category_repository.UpdateCategory(db, category)
		if err != nil {
			log.Printf("Erreur lors de la mise à jour de la catégorie: %v", err)
			http.Error(w, http.StatusText(500), 500) // Internal Server Error
			return
		}

		// Envoie d'une réponse réussie
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(categoryID)
	}
}


func DeleteCategoryHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")

        if id == "" {
			log.Printf("Erreur lors de la lecture de la requête: %v", id)
			http.Error(w, http.StatusText(400), 400) // Bad Request
			return
        }

        err := category_repository.DeleteCategory(db, id)
        if err != nil {
			log.Printf("Erreur lors de la suppression de la catégorie: %v", err)
			http.Error(w, http.StatusText(500), 500) // Internal Server Error
			return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Catégorie supprimée avec succès"))
    }
}
