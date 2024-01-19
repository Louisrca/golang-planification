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


func GetCustomerByID(db *sql.DB, id int) (model.Customer, error) {
	var u model.Customer
	err := db.QueryRow("SELECT id, firstname, email FROM customer WHERE id = ?", id).Scan(&u.ID, &u.FirstName, &u.Email)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return u, err
	}

	return u, nil
}

func CreateCustomer(db *sql.DB, customer model.Customer) (int64, error) {
    result, err := db.Exec("INSERT INTO customer (firstname, email) VALUES (?, ?)", customer.FirstName, customer.Email)
    if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        log.Printf("Erreur lors de la récupération de LastInsertId: %v", err)
        return 0, err
    }

    return id, nil
}

func UpdateCustomer(db *sql.DB, customer model.Customer) (int64, error) {
	result, err := db.Exec("UPDATE customer SET firstname = ?, email = ? WHERE id = ?", customer.FirstName, customer.Email, customer.ID)
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

func DeleteCustomer(db *sql.DB, customer_id int) error{
	query := `DELETE FROM category WHERE id = ?`
    _, err := db.Exec(query, customer_id)
    if err != nil {
        log.Printf("Erreur lors de la suppression de la catégorie: %v", err)
        return err
    }
    return nil
}