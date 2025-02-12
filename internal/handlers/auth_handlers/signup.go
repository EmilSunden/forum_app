package auth_handlers

import (
	"app/auth"
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Signup(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authRequest := models.AuthRequest{}
		// Create a variable to hold the input
		decoder := json.NewDecoder(r.Body)                   // Create a decoder to parse the JSON from the request body
		if err := decoder.Decode(&authRequest); err != nil { // Decode the JSON into the input variable
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Basic validation: ensurer both username and password are provided
		if authRequest.Username == "" || authRequest.Password == "" {
			http.Error(w, "Bad request: unable to parse JSON", http.StatusBadRequest)
			return
		}

		// Check if user already exists
		userExists, err := services.GetUserByUsername(db, authRequest.Username)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if userExists != nil {
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := auth.HashPassword(authRequest.Password)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Create the user
		user := models.User{
			Username: authRequest.Username,
			Password: hashedPassword,
		}

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User %s created successfully", authRequest.Username)

	}
}
