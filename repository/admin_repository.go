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