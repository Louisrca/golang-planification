package repository

import (
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func GetHairSalon(db *sql.DB) ([]model.HairSalon, error) {
	rows, err := db.Query("SELECT id, name, address, description, is_accepted FROM hair_salon")
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return nil, err
	}
	defer rows.Close()

	var hairSalons []model.HairSalon
	for rows.Next() {
		var u model.HairSalon
		if err := rows.Scan(&u.ID, &u.Name, &u.Address, &u.Description, &u.IsAccepted); err != nil {
			return nil, err
		}
		hairSalons = append(hairSalons, u)
	}

	return hairSalons, nil
}

func GetHairSalonByID(db *sql.DB, id string) (model.HairSalon, error) {
	var hairSalon model.HairSalon
	err := db.QueryRow("SELECT id, name, address, description, is_accepted FROM hair_salon WHERE id = ?", id).Scan(&hairSalon.ID, &hairSalon.Name, &hairSalon.Address, &hairSalon.Description, &hairSalon.IsAccepted)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return hairSalon, err
	}

	return hairSalon, nil
}

func CreateHairSalon(db *sql.DB, hairSalon model.HairSalon) (model.HairSalon, error) {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO hair_salon (id, name, address, description, is_accepted) VALUES (?, ?, ?, ?, ?)", uuid.String(), hairSalon.Name, hairSalon.Address, hairSalon.Description, hairSalon.IsAccepted)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.HairSalon{}, err
	}
	hairSalon.ID = uuid.String()

	return hairSalon, nil
}

func UpdateHairSalon(db *sql.DB, hairSalon model.HairSalon) (model.HairSalon, error) {
	_, err := db.Exec("UPDATE hair_salon SET name = ?, address = ?, description = ?, is_accepted = ? WHERE id = ?", hairSalon.Name, hairSalon.Address, hairSalon.Description, hairSalon.IsAccepted, hairSalon.ID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.HairSalon{}, err
	}

	var updatedHairSalon model.HairSalon
	err = db.QueryRow("SELECT id, name, address, description, is_accepted FROM hair_salon WHERE id = ?", hairSalon.ID).Scan(&updatedHairSalon.ID, &updatedHairSalon.Name, &updatedHairSalon.Address, &updatedHairSalon.Description, &updatedHairSalon.IsAccepted)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.HairSalon{}, err
	}

	return updatedHairSalon, nil
}

func DeleteHairSalon(db *sql.DB, id string) (model.HairSalon, error) {
	var hairSalon model.HairSalon
	err := db.QueryRow("DELETE FROM hair_salon WHERE id = ?", id).Scan(&hairSalon.ID, &hairSalon.Name, &hairSalon.Address, &hairSalon.Description, &hairSalon.IsAccepted)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.HairSalon{}, nil
		}
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.HairSalon{}, err
	}

	_, err = db.Exec("DELETE FROM hair_salon WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.HairSalon{}, err
	}

	return hairSalon, nil
}

func AcceptHairSalon(db *sql.DB, id string) (model.HairSalon, error) {
	_, err := db.Exec("UPDATE hair_salon SET is_accepted = TRUE WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du salon de coiffure: %v", err)
		return model.HairSalon{}, err
	}

	var updatedHairSalon model.HairSalon
	err = db.QueryRow("SELECT id, name, address, description, is_accepted FROM hair_salon WHERE id = ?", id).Scan(&updatedHairSalon.ID, &updatedHairSalon.Name, &updatedHairSalon.Address, &updatedHairSalon.Description, &updatedHairSalon.IsAccepted)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.HairSalon{}, err
	}

	return updatedHairSalon, nil
}
