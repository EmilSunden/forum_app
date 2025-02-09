package controllers

import (
	"encoding/json"
	"net/http"
)

func LogoutController(w http.ResponseWriter, r *http.Request) {
	// LogoutController is the function that handles the logout logic for the application
	// Logout logic here
	response := map[string]string{"message": "This is the logout endpoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.Write([]byte("LogoutController"))
}
