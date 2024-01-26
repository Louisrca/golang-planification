package controller

import (
	"api-planning/internal/utils"
	"api-planning/model"
	category_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FetchCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category, err := category_repository.GetCategory(db)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération des salon: %v", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(category)
	}
}

func FetchCategoryById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			utils.HandleError(w, "ID manquant dans l'URL", nil, http.StatusInternalServerError)
			return
		}

		category, err := category_repository.GetCategoryByID(db, id)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération de la catégorie: %v", err, http.StatusInternalServerError)
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
			utils.HandleError(w, "Requête invalide", err, http.StatusInternalServerError)
			return
		}

		categoryID, err := category_repository.CreateCategory(db, category)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la création de la categorie: %v", err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categoryID)
	}
}

func UpdateCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			utils.HandleError(w, "ID manquant dans l'URL", nil, http.StatusInternalServerError)
			return
		}

		var category model.Category

		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la récupération de la catégorie", err, http.StatusInternalServerError)
			return
		}

		category.ID = id

		updatedCategory, err := category_repository.UpdateCategory(db, category)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la mise à jour de la catégorie: %v", err, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updatedCategory)
	}
}

func DeleteCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			utils.HandleError(w, "ID manquant dans l'URL", nil, http.StatusInternalServerError)
			return
		}

		_, err := category_repository.DeleteCategory(db, id)
		if err != nil {
			utils.HandleError(w, "Erreur lors de la suppression de la catégorie: %v", err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Catégorie supprimée avec succès"))
	}
}
