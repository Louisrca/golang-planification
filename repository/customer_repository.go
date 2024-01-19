package repository

import (
	"api-planning/model"
	"database/sql"
	"log"
)


func GetCustomer(db *sql.DB) ([]model.Customer, error) {
    rows, err := db.Query("SELECT id, firstname, email FROM customer")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var customers []model.Customer
    for rows.Next() {
        var u model.Customer
        if err := rows.Scan(&u.ID, &u.FirstName, &u.Email); err != nil {
            return nil, err
        }
        customers = append(customers, u)
    }

    return customers, nil
}