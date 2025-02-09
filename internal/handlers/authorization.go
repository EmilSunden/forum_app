package handlers

import (
	"app/internal/controllers"
	"net/http"
)

func AuthMux() http.Handler {
	// AuthMux is the function that contains the authorization logic for the application
	// Auth logic here
	authMux := http.NewServeMux()
	authMux.Handle("/signup", Post(http.HandlerFunc(controllers.SignupController)))
	authMux.Handle("/login", http.HandlerFunc(controllers.LoginController))
	authMux.Handle("/logout", http.HandlerFunc(controllers.LogoutController))

	return http.StripPrefix("/api/v1/auth", authMux)
}
