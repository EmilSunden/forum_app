package models

import "time"

// BaseMessage holds fields common to both DMs and MediaPosts.
type BaseMessage struct {
	ID        int       `json:"id"`
	SenderID  int       `json:"sender_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// DirectMessage represent a private conversation between two users.
type DirectMessage struct {
	BaseMessage
	ReceiverID int `json:"receiver_id"`
}

// MediaPost represents a post/comment attached to a specific media item.
type MediaPost struct {
	BaseMessage
	MediaID int `json:"media_id"`
}
