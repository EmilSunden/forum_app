package handlers

import (
	"app/internal/handlers/auth_handlers"
	"net/http"

	"gorm.io/gorm"
)

func AuthMux(db *gorm.DB) http.Handler {
	// AuthMux is the function that contains the authorization logic for the application
	authMux := http.NewServeMux()
	authMux.Handle("/signup", auth_handlers.Signup(db))
	authMux.Handle("/login", auth_handlers.Login(db))
	authMux.Handle("/logout", auth_handlers.Logout())

	return http.StripPrefix("/api/v1/auth", authMux)
}
