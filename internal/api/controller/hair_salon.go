package controller

import (
	"api-planning/model"
	hair_salon_repository "api-planning/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)


func FetchHairSalon(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        hair_salons, err := hair_salon_repository.GetHairSalon(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des salon: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(hair_salons)
    }
}

func FetchHairSalonById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		hair_salon_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de hair_salon: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		hair_salon, err := hair_salon_repository.GetHairSalonByID(db,hair_salon_id)
		 if err != nil {
			log.Printf("Erreur lors de la récupération d: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hair_salon)
	}
}


func CreateHairSalonHandler(db *sql.DB) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {

		var hair_salon model.HairSalon

		err := json.NewDecoder(r.Body).Decode(&hair_salon)
		if err != nil {
			log.Printf("Erreur lors de la récupération de hair_salon: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		_, err = hair_salon_repository.CreateHairSalon(db, hair_salon)
		if err != nil {
			log.Printf("Erreur lors de la création de hair_salon: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hair_salon)
	}
}

func UpdateHairSalonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var hair_salon model.HairSalon

		err := json.NewDecoder(r.Body).Decode(&hair_salon)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		hair_salon_id, err := hair_salon_repository.UpdateHairSalon(db, hair_salon)
		if err != nil {
			log.Printf("Erreur lors de la création de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hair_salon_id)
	}
}

func DeleteHairSalonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")
		hair_salon_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'ID de hair_salon: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		err = hair_salon_repository.DeleteHairSalon(db, hair_salon_id)
		if err != nil {
			log.Printf("Erreur lors de la suppression de l'admin: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(hair_salon_id)
	}
}