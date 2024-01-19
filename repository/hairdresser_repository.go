package repository


import (
	"api-planning/model"
	"database/sql"
	"log"
)

func GetHairDresser(db *sql.DB) ([]model.Hairdresser, error) {
    rows, err := db.Query("SELECT id, firstname, email FROM hairdresser")
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

func GetHairDresserByID(db *sql.DB, id int) (model.Hairdresser, error) {
    var u model.Hairdresser
    err := db.QueryRow("SELECT id, firstname, email FROM hairdresser WHERE id = ?", id).Scan(&u.ID, &u.FirstName, &u.Email)
    if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return u, err
    }

    return u, nil
}

func CreateHairDresser(db *sql.DB, hairDresser model.Hairdresser) (int64, error) {
    result, err := db.Exec("INSERT INTO hairdresser (firstname, email) VALUES (?, ?)", hairDresser.FirstName, hairDresser.Email)
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


func UpdateHairDresser(db *sql.DB, hairDresser model.Hairdresser) (int64, error) {
    result, err := db.Exec("UPDATE hairdresser SET firstname = ?, email = ? WHERE id = ?", hairDresser.FirstName, hairDresser.Email, hairDresser.ID)
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

func DeleteHairDresser(db *sql.DB, id int) (int64, error) {
    result, err := db.Exec("DELETE FROM hairdresser WHERE id = ?", id)
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