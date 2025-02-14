package handlers

import (
	"app/internal/auth"
	"app/internal/models"
	"app/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Signup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authRequest models.AuthRequest
		if err := c.ShouldBindJSON(&authRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if authRequest.Username == "" || authRequest.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: unable to parse JSON"})
			return
		}

		userExists, err := services.GetUserByID(db, int64(authRequest.ID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		if userExists != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		hashedPassword, err := auth.HashPassword(authRequest.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		user := models.User{
			Username: authRequest.Username,
			Password: hashedPassword,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("User %s created successfully", authRequest.Username)})
	}
}
