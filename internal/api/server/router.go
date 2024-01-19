package router


import (
	"database/sql"
	"net/http"
	controller "api-planning/internal/api/controller"
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

	r.Get("/admin", controller.FetchAdmin(db))
	r.Get("/customer", controller.FetchCustomer(db))
	r.Get("/hair_salon", controller.FetchHairSalon(db))
	r.Get("/hairdresser", controller.FetchHairDresser(db))
	r.Get("/category", controller.FetchCategory(db))
	r.Get("/slot", controller.FetchSlot(db))
	r.Get("/booking", controller.FetchBooking(db))


    return r
}
