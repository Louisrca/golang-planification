package controller

import (
	"api-planning/model"
	category_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		id := r.URL.Query().Get("id")
		category_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de category: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		category, err := category_repository.GetCategoryByID(db,category_id)
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

		var category model.Category

		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			log.Printf("Erreur lors de la récupération de category: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		categoryID, err := category_repository.UpdateCategory(db, category)
		if err != nil {
			log.Printf("Erreur lors de la mise à jour de category: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(categoryID)
	}
}


func DeleteCategoryHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")
        categoryID, err := strconv.Atoi(id)
        if err != nil {
            http.Error(w, "ID de catégorie invalide", http.StatusBadRequest)
            return
        }

        err = category_repository.DeleteCategory(db, categoryID)
        if err != nil {
            http.Error(w, "Erreur lors de la suppression de la catégorie", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Catégorie supprimée avec succès"))
    }
}
