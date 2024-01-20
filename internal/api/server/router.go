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
	r.Put(("/admin/update"), controller.UpdateAdminHandler(db))
	r.Delete(("/admin/delete"), controller.DeleteAdminHandler(db))

	// booking routes
	r.Get("/booking", controller.FetchBooking(db))
	r.Post(("/booking/add"), controller.CreateBookingHandler(db))
	r.Put(("/booking/update"), controller.UpdateBookingHandler(db))
	r.Delete(("/booking/delete"), controller.DeleteBookingHandler(db))

	// category routes
	r.Get("/category", controller.FetchCategory(db))
	r.Post(("/category/add"), controller.CreateCategoryHandler(db))
	r.Put(("/category/update"), controller.UpdateCategoryHandler(db))
	r.Delete(("/category/delete"), controller.DeleteCategoryHandler(db))

	// customer routes
	r.Get("/customer", controller.FetchCustomer(db))
	r.Get("/customer/{id}", controller.FetchCustomerById(db))
	r.Post(("/customer/add"), controller.CreateCustomerHandler(db))
	r.Put(("/customer/update/{id}"), controller.UpdateCustomerHandler(db))
	r.Delete(("/customer/delete/{id}"), controller.DeleteCustomerHandler(db))

	// hair_salon routes
	r.Get("/hair_salon", controller.FetchHairSalon(db))
	r.Post(("/hair_salon/add"), controller.CreateHairSalonHandler(db))
	r.Put(("/hair_salon/update"), controller.UpdateHairSalonHandler(db))
	r.Delete(("/hair_salon/delete"), controller.DeleteHairSalonHandler(db))

	// hairdresser routes
	r.Get("/hairdresser", controller.FetchHairDresser(db))
	r.Post(("/hairdresser/add"), controller.CreateHairDresserHandler(db))
	r.Put(("/hairdresser/update"), controller.UpdateHairDresserHandler(db))
	r.Delete(("/hairdresser/delete"), controller.DeleteHairDresserHandler(db))

	// slog routes
	r.Get("/slot", controller.FetchSlot(db))
	r.Post(("/slot/add"), controller.CreateSlotHandler(db))
	r.Put(("/slot/update"), controller.UpdateSlotHandler(db))
	r.Delete(("/slot/delete"), controller.DeleteSlotHandler(db))

	// service routes
	r.Get("/service", controller.FetchService(db))
	r.Post(("/service/add"), controller.CreateServiceHandler(db))
	r.Put(("/service/update"), controller.UpdateServiceHandler(db))
	r.Delete(("/service/delete"), controller.DeleteServiceHandler(db))

	return r
}
