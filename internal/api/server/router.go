package router

import (
	controller "api-planning/internal/api/controller"
	"database/sql"
	"net/http"

	isAuthenticated "api-planning/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(db *sql.DB) *chi.Mux {
	
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Autoriser toutes les origines, ajustez selon vos besoins
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: true,
        MaxAge: 300, // Maximum cache duration for preflight requests
    })

    r.Use(cors.Handler)
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
	adminRoutes := chi.NewRouter()
	adminRoutes.Use(isAuthenticated.AuthMiddleware)
	adminRoutes.Get("/admin", controller.FetchAdmin(db))
	adminRoutes.Get("/admin/{id}", controller.FetchAdminById(db))
	adminRoutes.Post(("/admin/add"), controller.CreateAdminHandler(db))
	adminRoutes.Put(("/admin/update/{id}"), controller.UpdateAdminHandler(db))
	adminRoutes.Delete(("/admin/delete/{id}"), controller.DeleteAdminHandler(db))
	r.Mount("/", adminRoutes)

	// booking routes
	bookingRoutes := chi.NewRouter()
	bookingRoutes.Use(isAuthenticated.AuthMiddleware)
	bookingRoutes.Get("/booking", controller.FetchBooking(db))
	bookingRoutes.Get("/booking/{id}", controller.FetchBookingById(db))
	bookingRoutes.Post(("/booking/add"), controller.CreateBookingHandler(db))
	bookingRoutes.Put(("/booking/update/{id}"), controller.UpdateBookingHandler(db))
	bookingRoutes.Delete(("/booking/delete/{id}"), controller.DeleteBookingHandler(db))
	r.Mount("/", bookingRoutes)

	// category routes
	categoryRoutes := chi.NewRouter()
	categoryRoutes.Use(isAuthenticated.AuthMiddleware)
	categoryRoutes.Get("/category", controller.FetchCategory(db))
	categoryRoutes.Get("/category/{id}", controller.FetchCategoryById(db))
	categoryRoutes.Post(("/category/add"), controller.CreateCategoryHandler(db))
	categoryRoutes.Put(("/category/update/{id}"), controller.UpdateCategoryHandler(db))
	categoryRoutes.Delete(("/category/delete/{id}"), controller.DeleteCategoryHandler(db))
	r.Mount("/", categoryRoutes)

	// customer routes
	customerRoutes := chi.NewRouter()
	customerRoutes.Use(isAuthenticated.AuthMiddleware)
	customerRoutes.Get("/customer", controller.FetchCustomer(db))
	customerRoutes.Get("/customer/{id}", controller.FetchCustomerById(db))
	customerRoutes.Post(("/customer/add"), controller.CreateCustomerHandler(db))
	customerRoutes.Put(("/customer/update/{id}"), controller.UpdateCustomerHandler(db))
	customerRoutes.Delete(("/customer/delete/{id}"), controller.DeleteCustomerHandler(db))
	r.Mount("/", customerRoutes)

	// hair_salon routes
	hairSalonRoutes := chi.NewRouter()
	hairSalonRoutes.Use(isAuthenticated.AuthMiddleware)
	hairSalonRoutes.Get("/hair_salon", controller.FetchHairSalon(db))
	hairSalonRoutes.Get("/hair_salon/{id}", controller.FetchHairSalonById(db))
	hairSalonRoutes.Post(("/hair_salon/add"), controller.CreateHairSalonHandler(db))
	hairSalonRoutes.Put(("/hair_salon/update/{id}"), controller.UpdateHairSalonHandler(db))
	hairSalonRoutes.Delete(("/hair_salon/delete/{id}"), controller.DeleteHairSalonHandler(db))
	r.Mount("/", hairSalonRoutes)

	// hairdresser routes
	hairDresserRoutes := chi.NewRouter()
	hairDresserRoutes.Use(isAuthenticated.AuthMiddleware)
	hairDresserRoutes.Get("/hairdresser", controller.FetchHairDresser(db))
	hairDresserRoutes.Get("/hairdresser/{id}", controller.FetchHairDresserById(db))
	hairDresserRoutes.Post(("/hairdresser/add"), controller.CreateHairDresserHandler(db))
	hairDresserRoutes.Put(("/hairdresser/update/{id}"), controller.UpdateHairDresserHandler(db))
	hairDresserRoutes.Delete(("/hairdresser/delete/{id}"), controller.DeleteHairDresserHandler(db))
	r.Mount("/", hairDresserRoutes)

	// slog routes
	slogRoutes := chi.NewRouter()
	slogRoutes.Use(isAuthenticated.AuthMiddleware)
	slogRoutes.Get("/slot", controller.FetchSlot(db))
	slogRoutes.Get("/slot/{id}", controller.FetchSlotById(db))
	slogRoutes.Post(("/slot/add"), controller.CreateSlotHandler(db))
	slogRoutes.Put(("/slot/update/{id}"), controller.UpdateSlotHandler(db))
	slogRoutes.Delete(("/slot/delete/{id}"), controller.DeleteSlotHandler(db))
	r.Mount("/", slogRoutes)

	// service routes
	serviceRoutes := chi.NewRouter()
	serviceRoutes.Get("/service", controller.FetchService(db))
	serviceRoutes.Get("/service/{id}", controller.FetchServiceById(db))
	serviceRoutes.Post(("/service/add"), controller.CreateServiceHandler(db))
	serviceRoutes.Put(("/service/update/{id}"), controller.UpdateServiceHandler(db))
	serviceRoutes.Delete(("/service/delete/{id}"), controller.DeleteServiceHandler(db))
	r.Mount("/", serviceRoutes)

	return r
}
