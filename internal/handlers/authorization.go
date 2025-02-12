package handlers

import (
	"app/internal/controllers"
	"net/http"

	"gorm.io/gorm"
)

func AuthMux(db *gorm.DB) http.Handler {
	// AuthMux is the function that contains the authorization logic for the application
	// Auth logic here
	authMux := http.NewServeMux()
	authMux.Handle("/signup", Post(controllers.SignupController(db)))
	authMux.Handle("/login", Post(controllers.LoginController(db)))
	authMux.Handle("/logout", controllers.LogoutController())

	return http.StripPrefix("/api/v1/auth", authMux)
}
