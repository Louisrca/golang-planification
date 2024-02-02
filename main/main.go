package main

import (
	server "api-planning/internal/api/server"
	"api-planning/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connexion réussie à la base de données MariaDB !")
	router := server.NewRouter(db)

	// Démarrage du serveur
	log.Println("Démarrage du serveur sur le port :8080")
	http.ListenAndServe(":8080", router)
}
