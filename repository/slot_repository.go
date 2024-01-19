package repository

import (
	"api-planning/model"
	"database/sql"
	"log"
)


func GetSlot(db *sql.DB) ([]model.Slot, error) {
    rows, err := db.Query("SELECT id, hairdresser_id, is_booked, start_time, end_time FROM slot")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var slots []model.Slot
    for rows.Next() {
        var u model.Slot
        if err := rows.Scan(&u.ID, &u.HairdresserID,&u.IsBooked, &u.StartTime,&u.EndTime); err != nil {
            return nil, err
        }
        slots = append(slots, u)
    }

    return slots, nil
}


func GetSlotByID(db *sql.DB, id int) (model.Slot, error) {
	var u model.Slot
	err := db.QueryRow("SELECT id, hairdresser_id, is_booked, start_time, end_time FROM slot WHERE id = ?", id).Scan(&u.ID, &u.HairdresserID,&u.IsBooked, &u.StartTime,&u.EndTime)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return u, err
	}

	return u, nil
}


func CreateSlot(db *sql.DB, slot model.Slot) (int64, error) {
	result, err := db.Exec("INSERT INTO slot (hairdresser_id, is_booked, start_time, end_time) VALUES (?, ?, ?, ?)", slot.HairdresserID, slot.IsBooked, slot.StartTime, slot.EndTime)
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

func UpdateSlot(db *sql.DB, slot model.Slot) (int64, error) {
    query := `UPDATE slot SET start_time = ?, end_time = ?, is_booked = ?, hairdresser_id = ? WHERE id = ?`
    result, err := db.Exec(query, slot.StartTime, slot.EndTime, slot.IsBooked, slot.HairdresserID, slot.ID)
    if err != nil {
        log.Printf("Erreur lors de la mise à jour du slot: %v", err)
        return 0, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("Erreur lors de la récupération du nombre de lignes affectées: %v", err)
        return 0, err
    }

    return rowsAffected, nil
}


func DeleteSlot(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("DELETE FROM slot WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return 0, err
	}

	return result.RowsAffected()
}