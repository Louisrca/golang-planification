package repository


import (
	"api-planning/model"
	"database/sql"
	"log"
)

func GetAdmin(db *sql.DB) ([]model.Admin, error) {
    rows, err := db.Query("SELECT id, firstname, email FROM admin")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var admins []model.Admin
    for rows.Next() {
        var u model.Admin
        if err := rows.Scan(&u.ID, &u.FirstName, &u.Email); err != nil {
            return nil, err
        }
        admins = append(admins, u)
    }

    return admins, nil
}

func GetAdminById(db *sql.DB, id int) (model.Admin, error) {
	var admin model.Admin
	err := db.QueryRow("SELECT id, firstname, email FROM admin WHERE id = ?", id).Scan(&admin.ID, &admin.FirstName, &admin.Email)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return admin, err
	}

	return admin, nil
}

func CreateAdmin(db *sql.DB, admin model.Admin) (int64, error) {
	result, err := db.Exec("INSERT INTO admin (firstname, email) VALUES (?, ?)", admin.FirstName, admin.Email)
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