package services

import (
	"app/internal/models"
	"errors"
)

// HandleFriendRequest handles the business logic for accepting or rejecting a friend request.
func (s *FriendRequestService) HandleFriendRequest(requesterID int64, input models.FriendRequest) error {
	// Check if the friend request exists
	friendRequest, err := s.Data.HandleFriendRequest(requesterID, input.ReceiverID)
	if err != nil {
		return err
	}
	if friendRequest == nil {
		return errors.New("friend request not found")
	}

	if input.Status == "declined" {
		// If the user declined the request, delete the friend request
		s.Data.Notify(requesterID, "Your request was declined.")
		return s.Data.Delete(friendRequest)
	} else if input.Status == "accepted" {
		// Send acceptance notification to the requester
		s.Data.Notify(requesterID, "Your request was accepted.")
		// Create a friendship entry
		s.Data.CreateFriendshipEntry(requesterID, input.ReceiverID)
		// Delete the friend request entry
		return s.Data.Delete(friendRequest)
	}

	// Update the friend request status
	friendRequest.Status = input.Status
	friendRequest.Message = input.Message
	return s.Data.Update(friendRequest)
}
