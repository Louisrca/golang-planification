package repository

import (
	"api-planning/internal/utils"
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func GetHairDresser(db *sql.DB) ([]model.Hairdresser, error) {
	rows, err := db.Query("SELECT id, firstname, lastname, email, start_time, end_time, hair_salon_id FROM hairdresser")
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return nil, err
	}
	defer rows.Close()

	var hairdressers []model.Hairdresser
	for rows.Next() {
		var u model.Hairdresser
		if err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.StartTime, &u.EndTime, &u.HairSalonID); err != nil {
			return nil, err
		}
		hairdressers = append(hairdressers, u)
	}

	return hairdressers, nil
}

func GetHairDresserByID(db *sql.DB, id string) (model.Hairdresser, error) {
	var hairdresser model.Hairdresser
	err := db.QueryRow("SELECT id, firstname, lastname, email, start_time, end_time, hair_salon_id FROM hairdresser WHERE id = ?", id).Scan(&hairdresser.ID, &hairdresser.Firstname, &hairdresser.Lastname, &hairdresser.Email, &hairdresser.StartTime, &hairdresser.EndTime, &hairdresser.HairSalonID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return hairdresser, err
	}

	return hairdresser, nil
}

func CreateHairDresser(db *sql.DB, hairDresser model.Hairdresser) (model.Hairdresser, error) {
	uuid := uuid.New()

	hashedPassword := utils.HashPassword(hairDresser.Password)
	hairDresser.Password = hashedPassword
	_, err := db.Exec("INSERT INTO hairdresser (id, firstname, lastname, email, password, start_time, end_time, hair_salon_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", uuid.String(), hairDresser.Firstname, hairDresser.Lastname, hairDresser.Email, hairDresser.Password, hairDresser.StartTime, hairDresser.EndTime, hairDresser.HairSalonID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Hairdresser{}, err
	}

	hairDresser.ID = uuid.String()

	return hairDresser, nil
}

func UpdateHairDresser(db *sql.DB, hairDresser model.Hairdresser) (model.Hairdresser, error) {
	_, err := db.Exec("UPDATE hairdresser SET firstname = ?, lastname = ?, email = ?, start_time = ?, end_time = ?, hair_salon_id = ? WHERE id = ?", hairDresser.Firstname, hairDresser.Lastname, hairDresser.Email, hairDresser.StartTime, hairDresser.EndTime, hairDresser.HairSalonID, hairDresser.ID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Hairdresser{}, err
	}

	var updatedHairDresser model.Hairdresser
	err = db.QueryRow("SELECT id, firstname, lastname, email, start_time, end_time, hair_salon_id FROM hairdresser WHERE id = ?", hairDresser.ID).Scan(&updatedHairDresser.ID, &updatedHairDresser.Firstname, &updatedHairDresser.Lastname, &updatedHairDresser.Email, &updatedHairDresser.StartTime, &updatedHairDresser.EndTime, &updatedHairDresser.HairSalonID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Hairdresser{}, err
	}

	return updatedHairDresser, nil
}

func DeleteHairDresser(db *sql.DB, id string) (model.Hairdresser, error) {
	var hairDresser model.Hairdresser
	err := db.QueryRow("DELETE FROM hairdresser WHERE id = ?", id).Scan(&hairDresser.ID, &hairDresser.Firstname, &hairDresser.Lastname, &hairDresser.Email, &hairDresser.StartTime, &hairDresser.EndTime, &hairDresser.HairSalonID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Hairdresser{}, nil
		}
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Hairdresser{}, err
	}

	_, err = db.Exec("DELETE FROM hairdresser WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Hairdresser{}, err
	}

	return hairDresser, nil
}
