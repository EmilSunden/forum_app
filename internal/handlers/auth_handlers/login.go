package auth_handlers

import (
	"app/auth"
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func Login(db *gorm.DB) http.HandlerFunc {
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
		// Get user by username
		user, err := services.GetUserByUsername(db, authRequest.Username)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		// Validate the password
		if !auth.ValidatePassword(user.Password, authRequest.Password) {
			http.Error(w, "Invalid password", http.StatusBadRequest)
			return
		}

		// Generate a JWT token for the user
		token, err := auth.GenerateJWT(user.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			MaxAge:   int((time.Hour * 24).Seconds()),
		})
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User logged in successfully"})

	}
}
