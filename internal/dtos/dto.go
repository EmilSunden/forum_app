package dtos

// CreateFriendRequestInput represents the expected payload for creating a friend request.
type CreateFriendRequestInput struct {
	RequesterID int64  `json:"requester_id"`
	ReceiverID  int64  `json:"receiver_id"`
	Message     string `json:"message"`
}
