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


func GetServiceByID(db *sql.DB, id int) (model.Service, error) {
	var service model.Service
	err := db.QueryRow("SELECT id, name, price FROM service WHERE id = ?", id).Scan(&service.ID, &service.Name, &service.Price)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return service, err
	}

	return service, nil
}

func CreateService(db *sql.DB, service model.Service) (int64, error) {
	result, err := db.Exec("INSERT INTO service (name, price) VALUES (?, ?)", service.Name, service.Price)
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

func UpdateService(db *sql.DB, service model.Service) (int64, error) {
	result, err := db.Exec("UPDATE service SET name = ?, price = ? WHERE id = ?", service.Name, service.Price, service.ID)
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


func DeleteService(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("DELETE FROM service WHERE id = ?", id)
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