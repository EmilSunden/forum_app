package services

import "app/internal/dal"

// FriendRequestService handles friend request-related operations
type FriendRequestService struct {
	Data *dal.FriendRequestDataLayer
}
