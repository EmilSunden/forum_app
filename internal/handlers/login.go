package handlers

import (
	"app/internal/auth"
	"app/internal/models"
	"app/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authRequest := models.AuthRequest{}
		// Bind the JSON from the request body to the input variable
		if err := c.ShouldBindJSON(&authRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Basic validation: ensure both username and password are provided
		if authRequest.Username == "" || authRequest.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: unable to parse JSON"})
			return
		}

		// Get user by username
		user, err := services.GetUserByUsername(db, authRequest.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
			return
		}

		// Validate the password
		if !auth.ValidatePassword(user.Password, authRequest.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
			return
		}

		// Generate a JWT token for the user
		token, err := auth.GenerateJWT(int64(user.ID), user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.SetCookie("token", token, int((time.Hour * 24).Seconds()), "/", "", true, true)
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
	}
}
