package repository

import (
	"api-planning/internal/utils"
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func GetCustomer(db *sql.DB) ([]model.Customer, error) {
	rows, err := db.Query("SELECT id, firstname, lastname, email FROM customer")
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return nil, err
	}
	defer rows.Close()

	var customers []model.Customer
	for rows.Next() {
		var u model.Customer
		if err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email); err != nil {

			return nil, err
		}
		customers = append(customers, u)
	}

	return customers, nil
}

func GetCustomerByID(db *sql.DB, id string) (model.Customer, error) {
	var customer model.Customer
	err := db.QueryRow("SELECT id, firstname, lastname, email FROM customer WHERE id = ?", id).Scan(&customer.ID, &customer.Firstname, &customer.Lastname, &customer.Email)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return customer, err
	}

	return customer, nil
}

func CreateCustomer(db *sql.DB, customer model.Customer) (model.Customer, error) {
	uuid := uuid.New()

	hashedPassword := utils.HashPassword(customer.Password)
	customer.Password = hashedPassword

	_, err := db.Exec("INSERT INTO customer (id, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?)", uuid.String(), customer.Firstname, customer.Lastname, customer.Email, hashedPassword)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Customer{}, err
	}

	customer.ID = uuid.String()
	return customer, nil
}

func UpdateCustomer(db *sql.DB, customer model.Customer) (model.Customer, error) {
	_, err := db.Exec("UPDATE customer SET firstname = ?, lastname = ?, email = ? WHERE id = ?", customer.Firstname, customer.Lastname, customer.Email, customer.ID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Customer{}, err
	}

	var updatedCustomer model.Customer
	err = db.QueryRow("SELECT id, firstname, lastname, email FROM customer WHERE id = ?", customer.ID).Scan(&updatedCustomer.ID, &updatedCustomer.Firstname, &updatedCustomer.Lastname, &updatedCustomer.Email)
	if err != nil {
		log.Printf("Erreur lors de l'éxécution de la requête: %v", err)
		return model.Customer{}, err
	}

	return updatedCustomer, nil
}

func DeleteCustomer(db *sql.DB, id string) (model.Customer, error) {
	var customer model.Customer
	err := db.QueryRow("DELETE FROM customer WHERE id = ?", id).Scan(&customer.ID, &customer.Firstname, &customer.Lastname, &customer.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Customer{}, nil
		}
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Customer{}, err
	}

	_, err = db.Exec("DELETE FROM customer WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Customer{}, err
	}
	return customer, nil
}
