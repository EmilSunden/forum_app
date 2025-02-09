package handlers

import (
	"encoding/json"
	"net/http"
)

// Handlers is the function that contains all the handlers for the application
func PrivateHandler(w http.ResponseWriter, r *http.Request) {
	// Private handler
	response := map[string]string{"message": "This is a public endpoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	// Public handler
	response := map[string]string{"message": "This is a private endpoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
