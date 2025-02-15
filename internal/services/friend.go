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

// HandleFriendRequest handles the business logic for accepting or rejecting a friend request.
func (s *FriendRequestService) HandleFriendRequest(requesterID int64, input models.FriendRequest) error {
	// Example: Check if the friend request exists
	friendRequest, err := s.Repo.HandleFriendRequest(requesterID, input.ReceiverID)
	if err != nil {
		return err
	}
	if friendRequest == nil {
		return errors.New("friend request not found")
	}

	if input.Status == "declined" {
		// If the user declined the request, delete the friend request
		s.Repo.Notify(requesterID, "Your request was declined.")
		return s.Repo.Delete(friendRequest)
	} else if input.Status == "accepted" {
		// Send acceptance notification to the requester
		s.Repo.Notify(requesterID, "Your request was accepted.")
		// Create a friendship entry
		s.Repo.CreateFriendshipEntry(requesterID, input.ReceiverID)
		// Delete the friend request entry
		return s.Repo.Delete(friendRequest)
	}

	// Update the friend request status
	friendRequest.Status = input.Status
	friendRequest.Message = input.Message
	return s.Repo.Update(friendRequest)
}
