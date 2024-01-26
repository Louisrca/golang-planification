package repository

import (
	"api-planning/model"
	"database/sql"
	"log"

	"github.com/google/uuid"
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
		if err := rows.Scan(&u.ID, &u.HairdresserID, &u.IsBooked, &u.StartTime, &u.EndTime); err != nil {
			return nil, err
		}
		slots = append(slots, u)
	}

	return slots, nil
}

func GetSlotByID(db *sql.DB, id string) (model.Slot, error) {
	var u model.Slot
	err := db.QueryRow("SELECT id, hairdresser_id, is_booked, start_time, end_time FROM slot WHERE id = ?", id).Scan(&u.ID, &u.HairdresserID, &u.IsBooked, &u.StartTime, &u.EndTime)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return u, err
	}

	return u, nil
}

func CreateSlot(db *sql.DB, slot model.Slot) (model.Slot, error) {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO slot (id, hairdresser_id, is_booked, start_time, end_time) VALUES (?, ?, ?, ?, ?)",
		uuid.String(), slot.HairdresserID, slot.IsBooked, slot.StartTime, slot.EndTime)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête d'insertion: %v", err)
		return model.Slot{}, err
	}

	slot.ID = uuid.String()

	return slot, nil
}

func UpdateSlot(db *sql.DB, slot model.Slot) (model.Slot, error) {
	_, err := db.Exec("UPDATE slot SET hairdresser_id = ?, is_booked = ?, start_time = ?, end_time = ? WHERE id = ?",
		slot.HairdresserID, slot.IsBooked, slot.StartTime, slot.EndTime, slot.ID)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour du slot: %v", err)
		return model.Slot{}, err
	}
	var updatedSlot model.Slot
	err = db.QueryRow("SELECT id, start_time, end_time, is_booked, hairdresser_id FROM slot WHERE id = ?", slot.ID).
		Scan(&updatedSlot.ID, &updatedSlot.StartTime, &updatedSlot.EndTime, &updatedSlot.IsBooked, &updatedSlot.HairdresserID)
	if err != nil {
		log.Printf("Erreur lors de la récupération des informations mises à jour: %v", err)
		return model.Slot{}, err
	}

	return updatedSlot, nil
}

func DeleteSlot(db *sql.DB, id string) (model.Slot, error) {
	var slot model.Slot
	err := db.QueryRow("SELECT id, start_time, end_time, is_booked, hairdresser_id FROM slot WHERE id = ?", id).
		Scan(&slot.ID, &slot.StartTime, &slot.EndTime, &slot.IsBooked, &slot.HairdresserID)
	if err != nil {
		if err == sql.ErrNoRows {

			return model.Slot{}, nil
		}
		log.Printf("Erreur lors de la récupération du slot: %v", err)
		return model.Slot{}, err
	}

	_, err = db.Exec("DELETE FROM slot WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête de suppression: %v", err)
		return model.Slot{}, err
	}

	return slot, nil
}
