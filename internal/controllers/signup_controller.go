package controllers

import (
	"encoding/json"
	"net/http"
)

func SignupController(w http.ResponseWriter, r *http.Request) {
	// SignupController is the function that handles the signup logic for the application
	// Signup logic here
	response := map[string]string{"message": "This is the signup endpoint"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.Write([]byte("Signupcontroller"))
}
