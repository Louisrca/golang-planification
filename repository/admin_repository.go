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
        if err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email); err != nil {
            return nil, err
        }
        admins = append(admins, u)
    }

    return admins, nil
}

func GetAdminById(db *sql.DB, id string) (model.Admin, error) {
	var admin model.Admin
	err := db.QueryRow("SELECT id, firstname, lastname, email FROM admin WHERE id = ?", id).Scan(&admin.ID, &admin.Firstname, &admin.Lastname, &admin.Email)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return admin, err
	}

	return admin, nil
}

func CreateAdmin(db *sql.DB, admin model.Admin) (model.Admin, error) {

    uuid := uuid.New()

    _, err := db.Exec("INSERT INTO admin (id, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?)", 
                      uuid.String(), admin.Firstname, admin.Lastname, admin.Email, admin.Password)
    if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return model.Admin{}, err
    }

    return admin, nil
}



func UpdateAdmin(db *sql.DB, admin model.Admin) (model.Admin, error) {
    _, err := db.Exec("UPDATE admin SET firstname = ?, lastname = ?, email = ?, password = ? WHERE id = ?",
        admin.Firstname, admin.Lastname, admin.Email, admin.Password, admin.ID)
    if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête de mise à jour: %v", err)
        return model.Admin{}, err
    }

    var updatedAdmin model.Admin
    err = db.QueryRow("SELECT id, firstname, lastname, email, password  FROM admin WHERE id = ?", admin.ID).
        Scan(&updatedAdmin.ID, &updatedAdmin.Firstname, &updatedAdmin.Lastname, &updatedAdmin.Email, &updatedAdmin.Password)
    if err != nil {
        log.Printf("Erreur lors de la récupération des informations mises à jour: %v", err)
        return model.Admin{}, err
    }

    return updatedAdmin, nil
}





// func DeleteAdmin(db *sql.DB, id string) (model.Admin, error) {
// 	var admin model.Admin
// 	_, err := db.Exec("DELETE FROM admin WHERE id = ?", id)
// 	if err != nil {
// 		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
// 		return model.Admin{}, err
// 	}
	
// 	return admin, nil
// }



func DeleteAdmin(db *sql.DB, id string) (model.Admin, error) {
    var admin model.Admin
    err := db.QueryRow("DELETE from admin WHERE id= ?", id).Scan(&admin.ID, &admin.Firstname, &admin.Lastname, &admin.Email, &admin.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return model.Admin{}, nil
        }
        log.Printf("Erreur lors de la récupération du admin: %v", err)
        return model.Admin{}, err
    }

    _, err = db.Exec("DELETE FROM admin WHERE id = ?", id)
    if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête de suppression: %v", err)
        return model.Admin{}, err
    }

    return admin, nil
}