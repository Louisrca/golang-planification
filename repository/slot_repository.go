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