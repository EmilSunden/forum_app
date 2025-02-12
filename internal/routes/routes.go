package routes

import (
	"app/internal/handlers"
	"app/internal/middleware"
	"net/http"

	"gorm.io/gorm"
)

// Routes is the function that contains all the routes for the application
func Routes(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/auth/", handlers.AuthMux(db))

	protectedMux := middleware.AuthMiddleware(handlers.ProtectedMux())
	protectedGroup := http.StripPrefix("/api/v1/protected", protectedMux)
	mux.Handle("/api/v1/protected/", protectedGroup)

	return mux
}
