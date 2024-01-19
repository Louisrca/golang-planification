package repository

import (
	"api-planning/model"
	"database/sql"
	"log"
)


func GetService(db *sql.DB) ([]model.Service, error) {
    rows, err := db.Query("SELECT id, name, price FROM service")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var services []model.Service
    for rows.Next() {
        var u model.Service
        if err := rows.Scan(&u.ID, &u.Name, &u.Duration, &u.HairSalonID,&u.Price,&u.CategoryID); err != nil {
            return nil, err
        }
        services = append(services, u)
    }

    return services, nil
}