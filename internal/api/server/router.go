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
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
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
	adminRoutes.Get("/", controller.FetchAdmin(db))
	adminRoutes.Get("/{id}", controller.FetchAdminById(db))
	adminRoutes.Post(("/add"), controller.CreateAdminHandler(db))
	adminRoutes.Put(("/update/{id}"), controller.UpdateAdminHandler(db))
	adminRoutes.Delete(("/delete/{id}"), controller.DeleteAdminHandler(db))
	r.Mount("/admin", adminRoutes)

	// booking routes
	bookingRoutes := chi.NewRouter()
	bookingRoutes.Use(isAuthenticated.AuthMiddleware)
	bookingRoutes.Get("/", controller.FetchBooking(db))
	bookingRoutes.Get("/{id}", controller.FetchBookingById(db))
	bookingRoutes.Post(("/add"), controller.CreateBookingHandler(db))
	bookingRoutes.Put(("/update/{id}"), controller.UpdateBookingHandler(db))
	bookingRoutes.Delete(("/delete/{id}"), controller.DeleteBookingHandler(db))
	r.Mount("/booking", bookingRoutes)

	// category routes
	categoryRoutes := chi.NewRouter()
	categoryRoutes.Use(isAuthenticated.AuthMiddleware)
	categoryRoutes.Get("/", controller.FetchCategory(db))
	categoryRoutes.Get("/{id}", controller.FetchCategoryById(db))
	categoryRoutes.Post(("/add"), controller.CreateCategoryHandler(db))
	categoryRoutes.Put(("/{id}"), controller.UpdateCategoryHandler(db))
	categoryRoutes.Delete(("/delete/{id}"), controller.DeleteCategoryHandler(db))
	r.Mount("/category", categoryRoutes)

	// customer routes
	customerRoutes := chi.NewRouter()
	customerRoutes.Use(isAuthenticated.AuthMiddleware)
	customerRoutes.Get("/", controller.FetchCustomer(db))
	customerRoutes.Get("/{id}", controller.FetchCustomerById(db))
	customerRoutes.Post(("/add"), controller.CreateCustomerHandler(db))
	customerRoutes.Put(("/update/{id}"), controller.UpdateCustomerHandler(db))
	customerRoutes.Delete(("/delete/{id}"), controller.DeleteCustomerHandler(db))
	r.Mount("/customer", customerRoutes)

	// hair_salon routes
	hairSalonRoutes := chi.NewRouter()
	hairSalonRoutes.Use(isAuthenticated.AuthMiddleware)
	hairSalonRoutes.Get("/", controller.FetchHairSalon(db))
	hairSalonRoutes.Get("/{id}", controller.FetchHairSalonById(db))
	hairSalonRoutes.Post(("/add"), controller.CreateHairSalonHandler(db))
	hairSalonRoutes.Put(("/update/{id}"), controller.UpdateHairSalonHandler(db))
	hairSalonRoutes.Delete(("/delete/{id}"), controller.DeleteHairSalonHandler(db))
	r.Mount("/hair_salon", hairSalonRoutes)

	// hairdresser routes
	hairdresserRoutes := chi.NewRouter()
	hairdresserRoutes.Use(isAuthenticated.AuthMiddleware)
	hairdresserRoutes.Get("/", controller.FetchHairdresser(db))
	hairdresserRoutes.Get("/{id}", controller.FetchHairdresserById(db))
	hairdresserRoutes.Post(("/add"), controller.CreateHairdresserHandler(db))
	// hairdresserRoutes.Put(("/update/{id}"), controller.UpdateHairdresserHandler(db))
	hairdresserRoutes.Delete(("/delete/{id}"), controller.DeleteHairdresserHandler(db))
	r.Mount("/hairdresser", hairdresserRoutes)

	// slog routes
	slogRoutes := chi.NewRouter()
	slogRoutes.Use(isAuthenticated.AuthMiddleware)
	slogRoutes.Get("/", controller.FetchSlot(db))
	slogRoutes.Get("/{id}", controller.FetchSlotById(db))
	slogRoutes.Post(("/add"), controller.CreateSlotHandler(db))
	slogRoutes.Put(("/update/{id}"), controller.UpdateSlotHandler(db))
	slogRoutes.Delete(("/delete/{id}"), controller.DeleteSlotHandler(db))
	r.Mount("/slot", slogRoutes)

	// service routes
	serviceRoutes := chi.NewRouter()
	serviceRoutes.Get("/", controller.FetchService(db))
	serviceRoutes.Get("/{id}", controller.FetchServiceById(db))
	serviceRoutes.Post(("/add"), controller.CreateServiceHandler(db))
	serviceRoutes.Put(("/update/{id}"), controller.UpdateServiceHandler(db))
	serviceRoutes.Delete(("/delete/{id}"), controller.DeleteServiceHandler(db))
	r.Mount("/service", serviceRoutes)

	return r
}
