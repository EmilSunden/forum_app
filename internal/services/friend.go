package services

import (
	"app/internal/models"
	"app/internal/repositories"
	"errors"
)

// FriendRequestService handles friend request-related operations
type FriendRequestService struct {
	Repo *repositories.FriendRequestRepository
}

// CreateFriendRequest handles the business logic for creating a friend request.
func (s *FriendRequestService) CreateFriendRequest(requesterID int64, input models.FriendRequest) error {
	// Example: Check if a friend request already exists between these users.
	exists, err := s.Repo.Exists(requesterID, input.ReceiverID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("friend request already exists")
	}
	// Set initial status for the friend request
	input.Status = "pending"
	input.RequesterID = requesterID
	return s.Repo.Create(&input)
}
