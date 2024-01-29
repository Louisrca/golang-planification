package router

import (
	controller "api-planning/internal/api/controller"
	"database/sql"
	"net/http"

	isAuthenticated "api-planning/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Définir les routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenue sur notre serveur !"))
	})

	// register route
	r.Post("/register/customer", controller.RegisterCustomerHandler(db))
	r.Post("/register/admin", controller.RegisterAdminHandler(db))
	r.Post("/register/hairdresser", controller.RegisterHaidresserHandler(db))

	// login route
	r.Post("/login/admin", controller.LoginAdminHandler(db))
	r.Post("/login/customer", controller.LoginCustomerHandler(db))
	r.Post("/login/hairdresser", controller.LoginHairdresserHandler(db))
	

	// admin routes
	r.Get("/admin", controller.FetchAdmin(db))
	r.Get("/admin/{id}", controller.FetchAdminById(db))
	r.Post(("/admin/add"), controller.CreateAdminHandler(db))
	r.Put(("/admin/update/{id}"), controller.UpdateAdminHandler(db))
	r.Delete(("/admin/delete/{id}"), controller.DeleteAdminHandler(db))

	// booking routes
	r.Get("/booking", controller.FetchBooking(db))
	r.Get("/booking/{id}", controller.FetchBookingById(db))
	r.Post(("/booking/add"), controller.CreateBookingHandler(db))
	r.Put(("/booking/update/{id}"), controller.UpdateBookingHandler(db))
	r.Delete(("/booking/delete/{id}"), controller.DeleteBookingHandler(db))

	// category routes
	r.Get("/category", controller.FetchCategory(db))
	r.Get("/category/{id}", controller.FetchCategoryById(db))
	r.Post(("/category/add"), controller.CreateCategoryHandler(db))
	r.Put(("/category/update/{id}"), controller.UpdateCategoryHandler(db))
	r.Delete(("/category/delete/{id}"), controller.DeleteCategoryHandler(db))

	// customer routes
	customerRoutes := chi.NewRouter()
	customerRoutes.Use(isAuthenticated.AuthMiddleware)
	customerRoutes.Get("/customer", controller.FetchCustomer(db))
	customerRoutes.Get("/customer/{id}", controller.FetchCustomerById(db))
	customerRoutes.Post(("/customer/add"), controller.CreateCustomerHandler(db))
	customerRoutes.Put(("/customer/update/{id}"), controller.UpdateCustomerHandler(db))
	customerRoutes.Delete(("/customer/delete/{id}"), controller.DeleteCustomerHandler(db))
	r.Mount("/customer", customerRoutes)

	// hair_salon routes
	r.Get("/hair_salon", controller.FetchHairSalon(db))
	r.Get("/hair_salon/{id}", controller.FetchHairSalonById(db))
	r.Post(("/hair_salon/add"), controller.CreateHairSalonHandler(db))
	r.Put(("/hair_salon/update/{id}"), controller.UpdateHairSalonHandler(db))
	r.Delete(("/hair_salon/delete/{id}"), controller.DeleteHairSalonHandler(db))

	// hairdresser routes
	r.Get("/hairdresser", controller.FetchHairDresser(db))
	r.Get("/hairdresser/{id}", controller.FetchHairDresserById(db))
	r.Post(("/hairdresser/add"), controller.CreateHairDresserHandler(db))
	r.Put(("/hairdresser/update/{id}"), controller.UpdateHairDresserHandler(db))
	r.Delete(("/hairdresser/delete/{id}"), controller.DeleteHairDresserHandler(db))

	// slog routes
	r.Get("/slot", controller.FetchSlot(db))
	r.Get("/slot/{id}", controller.FetchSlotById(db))
	r.Post(("/slot/add"), controller.CreateSlotHandler(db))
	r.Put(("/slot/update/{id}"), controller.UpdateSlotHandler(db))
	r.Delete(("/slot/delete/{id}"), controller.DeleteSlotHandler(db))

	// service routes
	r.Get("/service", controller.FetchService(db))
	r.Get("/service/{id}", controller.FetchServiceById(db))
	r.Post(("/service/add"), controller.CreateServiceHandler(db))
	r.Put(("/service/update/{id}"), controller.UpdateServiceHandler(db))
	r.Delete(("/service/delete/{id}"), controller.DeleteServiceHandler(db))

	return r
}
