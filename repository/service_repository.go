package repository

import (
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func GetService(db *sql.DB) ([]model.Service, error) {
	rows, err := db.Query("SELECT id, name, price, duration,hair_salon_id,category_id  FROM service")
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return nil, err
	}
	defer rows.Close()

	var services []model.Service
	for rows.Next() {
		var u model.Service
		if err := rows.Scan(&u.ID, &u.Name, &u.Duration, &u.Price, &u.HairSalonID, &u.CategoryID); err != nil {
			return nil, err
		}
		services = append(services, u)
	}

	return services, nil
}

func GetServiceByID(db *sql.DB, id string) (model.Service, error) {
	var service model.Service
	err := db.QueryRow("SELECT id, name, price FROM service WHERE id = ?", id).Scan(&service.ID, &service.Name, &service.Price)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return service, err
	}

	return service, nil
}

func CreateService(db *sql.DB, service model.Service) (model.Service, error) {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO service (id, name, price, duration, category_id, hair_salon_id) VALUES (?, ?, ?, ?, ?, ?)",
		uuid.String(), service.Name, service.Price, service.Duration, service.CategoryID, service.HairSalonID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return model.Service{}, err
	}

	service.ID = uuid.String()

	return service, nil
}

func UpdateService(db *sql.DB, service model.Service) (model.Service, error) {
	_, err := db.Exec("UPDATE service SET name = ?, price = ?, duration = ?, category_id = ?, hair_salon_id = ? WHERE id = ?",
		service.Name, service.Price, service.Duration, service.CategoryID, service.HairSalonID, service.ID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête de mise à jour: %v", err)
		return model.Service{}, err
	}

	var updatedService model.Service
	err = db.QueryRow("SELECT id, name, price, duration, category_id, hair_salon_id FROM service WHERE id = ?", service.ID).
		Scan(&updatedService.ID, &updatedService.Name, &updatedService.Price, &updatedService.Duration, &updatedService.CategoryID, &updatedService.HairSalonID)
	if err != nil {
		log.Printf("Erreur lors de la récupération des informations mises à jour: %v", err)
		return model.Service{}, err
	}

	return updatedService, nil
}

func DeleteService(db *sql.DB, id string) (model.Service, error) {
	var service model.Service
	err := db.QueryRow("DELETE from service WHERE id= ?", id).Scan(&service.ID, &service.Name, &service.Price, &service.Duration, &service.CategoryID, &service.HairSalonID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Service{}, nil
		}
		log.Printf("Erreur lors de la récupération du service: %v", err)
		return model.Service{}, err
	}

	_, err = db.Exec("DELETE FROM service WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête de suppression: %v", err)
		return model.Service{}, err
	}

	return service, nil
}
