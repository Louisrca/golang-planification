package repository

import (
	"api-planning/model"
	"database/sql"
	"log"
)


func GetBooking(db *sql.DB) ([]model.Booking, error) {
    rows, err := db.Query("SELECT id, customer_id, service_id, slot_id, is_confirmed FROM booking")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var bookings []model.Booking
    for rows.Next() {
        var u model.Booking
        if err := rows.Scan(&u.ID, &u.CustomerID, &u.ServiceID,&u.SlotID,&u.IsConfirmed); err != nil {
            return nil, err
        }
        bookings = append(bookings, u)
    }

    return bookings, nil
}