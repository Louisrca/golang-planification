package repository

import (
	"api-planning/model"
	"database/sql"
	"log"
)


func GetHairSalon(db *sql.DB) ([]model.HairSalon, error) {
    rows, err := db.Query("SELECT id, name, address FROM hair_salon")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var hairSalons []model.HairSalon
    for rows.Next() {
        var u model.HairSalon
        if err := rows.Scan(&u.ID, &u.Name, &u.Address); err != nil {
            return nil, err
        }
        hairSalons = append(hairSalons, u)
    }

    return hairSalons, nil
}


func GetHairSalonByID(db *sql.DB, id int) (model.HairSalon, error) {
	var hairSalon model.HairSalon
	err := db.QueryRow("SELECT id, name, address FROM hair_salon WHERE id = ?", id).Scan(&hairSalon.ID, &hairSalon.Name, &hairSalon.Address)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return hairSalon, err
	}

	return hairSalon, nil
}


func CreateHairSalon(db *sql.DB, hairSalon model.HairSalon) (int64, error) {
	
	query := `INSERT INTO hair_salon (name, address) VALUES (?, ?)`
	result, err := db.Exec(query, hairSalon.Name, hairSalon.Address)
	if err != nil {
		log.Printf("Erreur lors de l'insertion de la réservation: %v", err)
	}
	id, err := result.LastInsertId()

	if err != nil {
		log.Printf("Erreur lors de la récupération de LastInsertId: %v", err)
		return 0, err

	}
	return id, nil
}

func UpdateHairSalon(db *sql.DB, hairSalon model.HairSalon) (int64, error) {
	
	query := `UPDATE hair_salon SET name = ?, address = ? WHERE id = ?`
	result, err := db.Exec(query, hairSalon.Name, hairSalon.Address, hairSalon.ID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erreur lors de la récupération de RowsAffected: %v", err)
		return 0, err
	}
	return rowsAffected, nil
}


func DeleteHairSalon(db *sql.DB, hairSalonID int) error{
	query:= `DELETE FROM hair_salon WHERE id = ?`
	_, err := db.Exec(query, hairSalonID)
	if err != nil {
		log.Printf("Erreur lors de la suppression du salon de coiffure: %v", err)
		return err
	}
	return nil
}