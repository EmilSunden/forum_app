package handlers

import (
	"app/internal/context"
	"app/internal/dal"
	"app/internal/dtos"
	"app/internal/models"
	"app/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFriendRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve the current user from the requests's context.
		user, ok := c.Get(string(context.UserKey))
		if !ok || user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Parse the incoming JSON payload into DTO
		var input dtos.CreateFriendRequestInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		data := dal.NewFriendRequest(db)
		friendService := services.FriendRequestService{Data: data}
		friendRequest := models.FriendRequest{
			ReceiverID: input.ReceiverID,
			Message:    input.Message,
		}

		// Use the service layer to create a friend request.
		if err := friendService.CreateFriendRequest(int64(user.(*models.User).ID), friendRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating friend request"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Friend request created successfully"})
	}
}
