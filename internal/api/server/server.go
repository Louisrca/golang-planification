package server

import (
	"api-planning/model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB) *chi.Mux {
    r := chi.NewRouter()

    // Middleware de base, vous pouvez ajouter le vôtre ici
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Définir les routes
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Bienvenue sur notre serveur !"))
    })

	r.Get("/users", usersHandler(db))


    return r
}

func getUsers(db *sql.DB) ([]model.Admin, error) {
    rows, err := db.Query("SELECT id, firstname, email FROM admin")
     if err != nil {
        log.Printf("Erreur lors de l'exécution de la requête: %v", err)
        return nil, err
    }
    defer rows.Close()

    var users []model.Admin
    for rows.Next() {
        var u model.Admin
        if err := rows.Scan(&u.ID, &u.FirstName, &u.Email); err != nil {
            return nil, err
        }
        users = append(users, u)
    }

    return users, nil
}


func usersHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        users, err := getUsers(db)
         if err != nil {
            log.Printf("Erreur lors de la récupération des utilisateurs: %v", err)
            http.Error(w, http.StatusText(500), 500)
            return
        }

        json.NewEncoder(w).Encode(users)
    }
}