package repository

import (
	"api-planning/model"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

func SendNotificationToAdmin(db *sql.DB, admin model.Admin, hairSalonID string) {
	message := "A new hair salon has been registered and is waiting for your approval."
	_, err := db.Exec("INSERT INTO notification (id, admin_id, message, hair_salon_id, is_read, created_at) VALUES (?, ?, ?, ?, FALSE, NOW())", uuid.NewString(), admin.ID, message, hairSalonID)

	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
	}
}

func GetNotificationByAdminID(db *sql.DB, adminID string) ([]model.Notification, error) {
	rows, err := db.Query("SELECT id, message, hair_salon_id, is_read, created_at FROM notification WHERE admin_id = ?", adminID)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
		return nil, err
	}
	defer rows.Close()

	var notifications []model.Notification
	for rows.Next() {
		var n model.Notification
		if err := rows.Scan(&n.ID, &n.Message, &n.HairSalonID, &n.IsRead, &n.CreatedAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}

	return notifications, nil
}
