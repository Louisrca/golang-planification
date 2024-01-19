package repository

import (
	"api-planning/model"
	"database/sql"
	"log"
)


func GetCategory(db *sql.DB) ([]model.Category, error) {
    rows, err := db.Query("SELECT id, name FROM category")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var category []model.Category
    for rows.Next() {
        var u model.Category
        if err := rows.Scan(&u.ID, &u.Name); err != nil {
            return nil, err
        }
        category = append(category, u)
    }

    return category, nil
}