package controller

import (
	hairdresser_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"api-planning/model"
)


func FetchHairDresser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        hairdressers, err := hairdresser_repository.GetHairDresser(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des salon: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(hairdressers)
    }
}

func FetchHairDresserById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		hairdresser_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de hairdresser: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		hairdresser, err := hairdresser_repository.GetHairDresserByID(db,hairdresser_id)
		 if err != nil {
			log.Printf("Erreur lors de la récupération d: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hairdresser)
	}
}

func CreateHairDresserHandler(db *sql.DB) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {

		var hairdresser model.Hairdresser

		err := json.NewDecoder(r.Body).Decode(&hairdresser)
		if err != nil {
			log.Printf("Erreur lors de la récupération de hairdresser: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		hairdresser_id, err := hairdresser_repository.CreateHairDresser(db, hairdresser)
		if err != nil {
			log.Printf("Erreur lors de la création de hairdresser: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hairdresser_id)
	}
}

func UpdateHairDresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var hairdresser model.Hairdresser

		err := json.NewDecoder(r.Body).Decode(&hairdresser)
		if err != nil {
			log.Printf("Erreur lors de la récupération de hairdresser: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		hairdresser_id, err := hairdresser_repository.UpdateHairDresser(db, hairdresser)
		if err != nil {
			log.Printf("Erreur lors de la création de hairdresser: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hairdresser_id)
	}
}

func DeleteHairDresserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		hairdresser_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de hairdresser: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		_,err = hairdresser_repository.DeleteHairDresser(db, hairdresser_id)
		if err != nil {
			log.Printf("Erreur lors de la suppression de hairdresser: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hairdresser_id)
	}
}