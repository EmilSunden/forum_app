package friends

import (
	"app/internal/context"
	"app/internal/users"

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
		var input CreateFriendRequestInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		data := NewFriendRequest(db)
		friendService := FriendRequestService{Data: data}
		friendRequest := FriendRequest{
			ReceiverID: input.ReceiverID,
			Message:    input.Message,
		}

		// Use the service layer to create a friend request.
		if err := friendService.CreateFriendRequest(int64(user.(*users.User).ID), friendRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating friend request"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Friend request created successfully"})
	}
}

func HandleFriendRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve the current user from the requests's context.
		user, ok := c.Get(string(context.UserKey))
		if !ok || user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// Parse the incoming friendrequest payload
		var friendrequest HandleFriendRequestInput
		if err := c.ShouldBindJSON(&friendrequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Create a new friend request data layer
		data := NewFriendRequest(db)
		friendService := FriendRequestService{Data: data}
		friendRequest := FriendRequest{
			RequesterID: friendrequest.RequesterID,
			ReceiverID:  friendrequest.ReceiverID,
			Status:      friendrequest.Status,
			Message:     friendrequest.Message,
		}
		// Use the service layer to handle the friend request
		if err := friendService.HandleFriendRequest(int64(user.(*users.User).ID), friendRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error handling friend request"})
			return
		}
		// Return a success message
		c.JSON(http.StatusOK, gin.H{"message": "Friend request handled successfully"})
	}
}
