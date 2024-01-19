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

func GetCategoryByID(db *sql.DB, id int) (model.Category, error) {
	var category model.Category
	err := db.QueryRow("SELECT id, name FROM category WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return category, err
	}

	return category, nil
}

func CreateCategory(db *sql.DB, category model.Category) (int64, error) {

	query := `INSERT INTO category (name) VALUES (?)`
	result, err := db.Exec(query, category.Name)
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

func UpdateCategory(db *sql.DB, category model.Category) (int64, error) {

	query := `UPDATE category SET name = ? WHERE id = ?`
	result, err := db.Exec(query, category.Name, category.ID)
	if err != nil {
		log.Printf("Erreur lors de l'insertion de la category: %v", err)
		return 0, err
	}
	id, err := result.RowsAffected()

	if err != nil {
		log.Printf("Erreur lors de la récupération de LastInsertId: %v", err)
	}
	return id, nil
}


func DeleteCategory(db *sql.DB, categoryID int) error {
    query := `DELETE FROM category WHERE id = ?`
    _, err := db.Exec(query, categoryID)
    if err != nil {
        log.Printf("Erreur lors de la suppression de la catégorie: %v", err)
        return err
    }
    return nil
}

