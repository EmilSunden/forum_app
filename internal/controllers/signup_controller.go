package controllers

import (
	"app/auth"
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type SignupInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignupController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// SignupController is the function that handles the signup logic for the application
		// Signup logic here
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

		var input SignupInput                          // Create a variable to hold the input
		decoder := json.NewDecoder(r.Body)             // Create a decoder to parse the JSON from the request body
		if err := decoder.Decode(&input); err != nil { // Decode the JSON into the input variable
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Basic validation: ensurer both username and password are provided
		if input.Username == "" || input.Password == "" {
			http.Error(w, "Bad request: unable to parse JSON", http.StatusBadRequest)
			return
		}

		// TODO: Add additional validation, passowrd hasing, and user creation logic here
		// For example, check if the user already exists, hash the password, and store the user in the database

		// Check if user already exists
		exists, err := services.UserExists(db, input.Username)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := auth.HashPassword(input.Password)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Create the user
		user := models.User{
			Username: input.Username,
			Password: hashedPassword,
		}

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User %s created successfully", input.Username)

	}
}
