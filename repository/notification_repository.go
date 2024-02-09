package repository

import (
	"api-planning/model"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

func SendNotificationToAdmin(db *sql.DB, admin model.Admin) {
	message := "A new hair salon has been registered and is waiting for your approval."
	_, err := db.Exec("INSERT INTO notification (id, admin_id, message, is_read, created_at) VALUES (?, ?, ?, FALSE, NOW())", uuid.NewString(), admin.ID, message)

	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête: %v", err)
	}
}
