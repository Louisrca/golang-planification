package repository

import (
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func GetAdmin(db *sql.DB) ([]model.Admin, error) {
    rows, err := db.Query("SELECT id, firstname, lastname, email FROM admin")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var admins []model.Admin
    for rows.Next() {
        var u model.Admin
        if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
            return nil, err
        }
        admins = append(admins, u)
    }

    return admins, nil
}

func GetAdminById(db *sql.DB, id int) (model.Admin, error) {
	var admin model.Admin
	err := db.QueryRow("SELECT id, firstname, lastname, email FROM admin WHERE id = ?", id).Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return admin, err
	}

	return admin, nil
}

func CreateAdmin(db *sql.DB, admin model.Admin) (model.Admin, error) {
    // Génération d'un nouveau UUID
    uuid := uuid.New()

    // Insertion de l'admin dans la base de données
    _, err := db.Exec("INSERT INTO admin (id, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?)", 
                      uuid.String(), admin.FirstName, admin.LastName, admin.Email, admin.Password)
    if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return model.Admin{}, err
    }
    // Mise à jour de l'ID de l'objet admin
    admin.ID = uuid.String()
    // Ici, vous pouvez choisir de retourner l'objet admin tel quel
    // ou de faire une requête supplémentaire pour récupérer toutes les données
    // de l'admin à partir de la base de données si nécessaire.

    return admin, nil
}




func UpdateAdmin(db *sql.DB, admin model.Admin) (int64, error) {
	result, err := db.Exec("UPDATE admin SET firstname = ?, email = ? WHERE id = ?", admin.FirstName, admin.Email, admin.ID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return 0, err
	}

	return result.RowsAffected()
}


func DeleteAdmin(db *sql.DB, id int) (int64, error) {
	
	result, err := db.Exec("DELETE FROM admin WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return 0, err
	}
	
	return result.RowsAffected()
}