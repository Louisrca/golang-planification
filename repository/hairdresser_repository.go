package repository


import (
	"api-planning/model"
	"database/sql"
	"log"
)

func GetHairDresser(db *sql.DB) ([]model.Hairdresser, error) {
    rows, err := db.Query("SELECT id, firstname, email FROM hair_dresser")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var hairDressers []model.Hairdresser
    for rows.Next() {
        var u model.Hairdresser
        if err := rows.Scan(&u.ID, &u.FirstName, &u.Email); err != nil {
            return nil, err
        }
        hairDressers = append(hairDressers, u)
    }

    return hairDressers, nil
}