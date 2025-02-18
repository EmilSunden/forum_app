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

func HandleFriendRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve the current user from the requests's context.
		user, ok := c.Get(string(context.UserKey))
		if !ok || user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// Parse the incoming friendrequest payload
		var friendrequest dtos.HandleFriendRequestInput
		if err := c.ShouldBindJSON(&friendrequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Create a new friend request data layer
		data := dal.NewFriendRequest(db)
		friendService := services.FriendRequestService{Data: data}
		friendRequest := models.FriendRequest{
			RequesterID: friendrequest.RequesterID,
			ReceiverID:  friendrequest.ReceiverID,
			Status:      friendrequest.Status,
			Message:     friendrequest.Message,
		}
		// Use the service layer to handle the friend request
		if err := friendService.HandleFriendRequest(int64(user.(*models.User).ID), friendRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error handling friend request"})
			return
		}
		// Return a success message
		c.JSON(http.StatusOK, gin.H{"message": "Friend request handled successfully"})
	}
}
