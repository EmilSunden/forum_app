package routes

import (
	"app/internal/handlers"
	"net/http"

	"app/internal/middleware"
)

// Routes is the function that contains all the routes for the application
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Public Routes
	mux.HandleFunc("/public", handlers.PublicHandler)

	// Private route wrapped with auth middleware
	privateHandler := middleware.AuthMiddleware(http.HandlerFunc(handlers.PrivateHandler))
	mux.Handle("/private", privateHandler)

	return mux
}
