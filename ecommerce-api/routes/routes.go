package routes

import (
	"net/http"

	"backend/ecommerce-api/handlers"
	"backend/ecommerce-api/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ecommerce API is running 🚀"))
	})
	r.HandleFunc("/api/auth/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/products", handlers.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/products", handlers.CreateProduct).Methods("POST") // could protect later

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/cart", handlers.AddToCart).Methods("POST")
	protected.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")

	return r
}
