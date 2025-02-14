package repositories

import (
	"app/internal/models"

	"gorm.io/gorm"
)

type FriendRequestRepository struct {
	db *gorm.DB
}

func NewFriendRequestRepository(db *gorm.DB) *FriendRequestRepository {
	return &FriendRequestRepository{db: db}
}

// Exists checks if a friend request already exists between two users.
func (r *FriendRequestRepository) Exists(requesterID, receiverID int64) (bool, error) {
	var count int64
	err := r.db.Model(&models.FriendRequest{}).Where("requester_id = ? And receiver_id = ?", requesterID, receiverID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Create creates a new friend request.
func (r *FriendRequestRepository) Create(fr *models.FriendRequest) error {
	return r.db.Create(fr).Error
}
