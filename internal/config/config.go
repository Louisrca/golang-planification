
package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

// InitDB initialise et retourne une connexion à la base de données.
func InitDB() *sql.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
    }

    dbUser := os.Getenv("MYSQL_USER")
    dbPassword := os.Getenv("MYSQL_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("MARIADB_DATABASE")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    return db
}
