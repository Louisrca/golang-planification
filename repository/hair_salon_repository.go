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