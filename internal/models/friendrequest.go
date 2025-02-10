package models

type FriendRequest struct {
	ID               int    `json:"id"`
	SenderUsername   string `json:"sender_username"`
	ReceiverUsername string `json:"receiver_username"`
}

// FriendRequest Status model

type FriendRequestStatus string

const (
	// Pending status
	Pending FriendRequestStatus = "pending"
	// Accepted status
	Accepted FriendRequestStatus = "accepted"
	// Rejected status
	Rejected FriendRequestStatus = "rejected"
)
