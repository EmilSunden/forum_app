package handlers

import (
	contextKeys "app/internal/contextkeys"
	"app/internal/dtos"
	"app/internal/models"
	"app/internal/repositories"
	"app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFriendRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve the current user from the requests's context.
		user, ok := c.Get(string(contextKeys.UserKey))
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

		repo := repositories.NewFriendRequestRepository(db)
		friendService := services.FriendRequestService{Repo: repo}
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

func HandleFriendRequest(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve the current user from the requests's context.
		user, ok := c.Get(string(contextKeys.UserKey))
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

		// Check if the friendrequest exists in the database
		repo := repositories.NewFriendRequestRepository(db)
		friendService := services.FriendRequestService{Repo: repo}
		friendRequest := models.FriendRequest{
			RequesterID: friendrequest.RequesterID,
			ReceiverID:  friendrequest.ReceiverID,
			Status:      friendrequest.Status,
			Message:     friendrequest.Message,
		}

		// Use the service layer to handle the friend request.
		if err := friendService.HandleFriendRequest(int64(user.(*models.User).ID), friendRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error handling friend request"})
			return
		}

		// If current user accepts the friend request create a friendshipentry in database and clean up the friendrequest entry from friendrequest table.

		// Send a notification to the requester that friend request has been accepted. Purge notification from database after sending it to the requester.

		// iIf current user declines the friend request, purge friendrequest entry from friendrequest table and send a notification to the requester that friendrequest has been declined.
		// Purge notification from database after sending it to the requester.
		c.JSON(http.StatusOK, gin.H{"message": "Friend request handled successfully"})
	}
}
