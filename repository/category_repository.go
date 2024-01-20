package repository

import (
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
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

func GetCategoryByID(db *sql.DB, id string) (model.Category, error) {
	var category model.Category
	err := db.QueryRow("SELECT id, name FROM category WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return category, err
	}

	return category, nil
}

func CreateCategory(db *sql.DB, category model.Category) (model.Category, error) {

	uuid := uuid.New()
	_, err := db.Exec(`INSERT INTO  category (id, name) VALUES (?,?)`, uuid.String(),category.Name)
	
	if err != nil {
		log.Printf("Erreur lors de l'insértion de la catégorie : %v", err)
		return model.Category{}, err
	}
	return category, nil
}

func UpdateCategory(db *sql.DB, category model.Category) (model.Category, error) {


	_, err := db.Exec(`UPDATE category SET name = ? WHERE id = ?`, category.Name, category.ID)
	if err != nil {
        log.Printf("Erreur lors de la mise à jour de la réservation: %v", err)
        return model.Category{}, err
    }
    return category, nil
}


func DeleteCategory(db *sql.DB, categoryID string) error {
    query := `DELETE FROM category WHERE id = ?`
    _, err := db.Exec(query, categoryID)
    if err != nil {
        log.Printf("Erreur lors de la suppression de la catégorie: %v", err)
        return err
    }
    return nil
}

