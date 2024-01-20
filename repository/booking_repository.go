package repository

import (
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
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

func GetBookingById(db *sql.DB, id string) (model.Booking, error) {
	var booking model.Booking
	err := db.QueryRow("SELECT id, customer_id, service_id, slot_id, is_confirmed FROM booking WHERE id = ?", id).Scan(&booking.ID, &booking.CustomerID, &booking.ServiceID,&booking.SlotID,&booking.IsConfirmed)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return booking, err
	}

	return booking, nil
}


func CreateBooking(db *sql.DB, booking model.Booking) (model.Booking, error) {

    uuid := uuid.New()
    _, err := db.Exec(`INSERT INTO booking (id, customer_id, service_id, slot_id, is_confirmed) VALUES (?, ?, ?, ?, ?)`, uuid.String(), booking.CustomerID, booking.ServiceID, booking.SlotID, booking.IsConfirmed)
    if err != nil {
        log.Printf("Erreur lors de l'insertion de la réservation: %v", err)
        return model.Booking{}, err
    }
    return booking, nil
}

func UpdateBooking(db *sql.DB, updatedBooking model.Booking) (model.Booking, error) {

    _, err := db.Exec(`UPDATE booking SET customer_id = ?, service_id = ?, slot_id = ?, is_confirmed = ? WHERE id = ?`, updatedBooking.CustomerID, updatedBooking.ServiceID, updatedBooking.SlotID, updatedBooking.IsConfirmed, updatedBooking.ID)
    if err != nil {
        log.Printf("Erreur lors de la mise à jour de la réservation: %v", err)
        return model.Booking{},err
    }
    return updatedBooking, nil
}



func DeleteBooking(db *sql.DB, bookingID string) error {
    query := `DELETE FROM booking WHERE id = ?`
    _, err := db.Exec(query, bookingID)
    if err != nil {
        log.Printf("Erreur lors de la suppression de la réservation: %v", err)
        return err
    }
    return nil
}