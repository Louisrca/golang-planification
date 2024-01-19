package router

import (
	controller "api-planning/internal/api/controller"
	"database/sql"
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

	// admin routes
	r.Get("/admin", controller.FetchAdmin(db))
	r.Post(("/admin/add"), controller.CreateAdminHandler(db))
	r.Post(("/admin/update"), controller.UpdateAdminHandler(db))
	r.Post(("/admin/delete"), controller.DeleteAdminHandler(db))

	// booking routes
	r.Get("/booking", controller.FetchBooking(db))
	r.Post(("/booking/add"), controller.CreateBookingHandler(db))
	r.Post(("/booking/update"), controller.UpdateBookingHandler(db))
	r.Post(("/booking/delete"), controller.DeleteBookingHandler(db))

	// category routes
	r.Get("/category", controller.FetchCategory(db))
	r.Post(("/category/add"), controller.CreateCategoryHandler(db))
	r.Post(("/category/update"), controller.UpdateCategoryHandler(db))
	r.Post(("/category/delete"), controller.DeleteCategoryHandler(db))


	// customer routes
	r.Get("/customer", controller.FetchCustomer(db))
	r.Post(("/customer/add"), controller.CreateCustomerHandler(db))
	r.Post(("/customer/update"), controller.UpdateCustomerHandler(db))
	r.Post(("/customer/delete"), controller.DeleteCustomerHandler(db))

	// hair_salon routes
	r.Get("/hair_salon", controller.FetchHairSalon(db))
	r.Post(("/hair_salon/add"), controller.CreateHairSalonHandler(db))
	r.Post(("/hair_salon/update"), controller.UpdateHairSalonHandler(db))
	r.Post(("/hair_salon/delete"), controller.DeleteHairSalonHandler(db))


	// hairdresser routes
	r.Get("/hairdresser", controller.FetchHairDresser(db))
	r.Post(("/hairdresser/add"), controller.CreateHairDresserHandler(db))
	r.Post(("/hairdresser/update"), controller.UpdateHairDresserHandler(db))
	r.Post(("/hairdresser/delete"), controller.DeleteHairDresserHandler(db))

	// slog routes
	r.Get("/slot", controller.FetchSlot(db))
	r.Post(("/slot/add"), controller.CreateSlotHandler(db))
	r.Post(("/slot/update"), controller.UpdateSlotHandler(db))
	r.Post(("/slot/delete"), controller.DeleteSlotHandler(db))

	// service routes
	r.Get("/service", controller.FetchService(db))
	r.Post(("/service/add"), controller.CreateServiceHandler(db))
	r.Post(("/service/update"), controller.UpdateServiceHandler(db))
	r.Post(("/service/delete"), controller.DeleteServiceHandler(db))

    return r
}
