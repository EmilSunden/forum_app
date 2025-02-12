package models

import "github.com/jinzhu/gorm"

// FriendRequest represents a pending (or resolved) friend request.
type FriendRequest struct {
	gorm.Model
	RequesterID int    `json:"requester_id"`
	ReceiverID  int    `json:"receiver_id"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}

// FriendshipEntry represents a unidirectional friendship link.
type FriendshipEntry struct {
	gorm.Model
	UserID       int `json:"user_id"`
	FriendUserID int `json:"friend_user_id"`
}

// FriendSummary represents a brief overview of a friend
type FriendSummary struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name,omitempty"`
	GivenName   string `json:"given_name,omitempty"`
	FamilyName  string `json:"family_name,omitempty"`
	Avatar      string `json:"avatar"`
}

// FriendList represents a user's aggregated list of friends.
type FriendList struct {
	UserID  int64           `json:"user_id"`
	Friends []FriendSummary `json:"friends"`
}
