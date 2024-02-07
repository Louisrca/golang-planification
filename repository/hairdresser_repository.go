package repository

import (
	"api-planning/internal/utils"
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func GetHairdresser(db *sql.DB) ([]model.Hairdresser, error) {
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

func GetHairdresserByID(db *sql.DB, id string) (model.Hairdresser, error) {
	var hairdresser model.Hairdresser
	err := db.QueryRow("SELECT id, firstname, lastname, email, start_time, end_time, hair_salon_id FROM hairdresser WHERE id = ?", id).Scan(&hairdresser.ID, &hairdresser.Firstname, &hairdresser.Lastname, &hairdresser.Email, &hairdresser.StartTime, &hairdresser.EndTime, &hairdresser.HairSalonID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return hairdresser, err
	}

	return hairdresser, nil
}

func GetHairdresserByEmail(db *sql.DB, email string) (model.Hairdresser, error) {
	var hairdresser model.Hairdresser
	err := db.QueryRow("SELECT email, password FROM hairdresser WHERE email = ?", email).Scan(&hairdresser.Email, &hairdresser.Password)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return hairdresser, err
	}

	return hairdresser, nil
}

func CreateHairdresser(db *sql.DB, hairdresser model.Hairdresser) (model.Hairdresser, error) {
	uuid := uuid.New()

	hashedPassword := utils.HashPassword(hairdresser.Password)
	hairdresser.Password = hashedPassword

	var hairSalonID sql.NullString
	if hairdresser.HairSalonID != "" {
		hairSalonID = sql.NullString{String: hairdresser.HairSalonID, Valid: true}
	} else {
		hairSalonID = sql.NullString{Valid: false}
	}

	_, err := db.Exec("INSERT INTO hairdresser (id, firstname, lastname, email, password, start_time, end_time, hair_salon_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", uuid.String(), hairdresser.Firstname, hairdresser.Lastname, hairdresser.Email, hairdresser.Password, hairdresser.StartTime, hairdresser.EndTime, hairSalonID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Hairdresser{}, err
	}

	hairdresser.ID = uuid.String()

	return hairdresser, nil
}


func UpdateHairdresser(db *sql.DB, hairdresserID string, hairSalonID string) (model.Hairdresser, error) {
    _, err := db.Exec("UPDATE hairdresser SET hair_salon_id = ? WHERE id = ?", hairSalonID, hairdresserID)
    if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return model.Hairdresser{}, err
    }

    var updatedHairdresser model.Hairdresser
    err = db.QueryRow("SELECT id, firstname, lastname, email, start_time, end_time, hair_salon_id FROM hairdresser WHERE id = ?", hairdresserID).Scan(&updatedHairdresser.ID, &updatedHairdresser.Firstname, &updatedHairdresser.Lastname, &updatedHairdresser.Email, &updatedHairdresser.StartTime, &updatedHairdresser.EndTime, &updatedHairdresser.HairSalonID)
    if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return model.Hairdresser{}, err
    }

    return updatedHairdresser, nil
}



func DeleteHairdresser(db *sql.DB, id string) (model.Hairdresser, error) {
	var hairdresser model.Hairdresser
	err := db.QueryRow("DELETE FROM hairdresser WHERE id = ?", id).Scan(&hairdresser.ID, &hairdresser.Firstname, &hairdresser.Lastname, &hairdresser.Email, &hairdresser.StartTime, &hairdresser.EndTime, &hairdresser.HairSalonID)
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

	return hairdresser, nil
}


