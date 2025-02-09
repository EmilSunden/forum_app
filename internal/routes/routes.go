package routes

import (
	"app/internal/handlers"
	"app/internal/middleware"
	"net/http"
)

// Routes is the function that contains all the routes for the application
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/auth/", handlers.AuthMux())

	protectedMux := middleware.AuthMiddleware(handlers.ProtectedMux())
	protectedGroup := http.StripPrefix("/api/v1/protected", protectedMux)
	mux.Handle("/api/v1/protected/", protectedGroup)

	return mux
}
