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

// HandleFriendRequest handles a friendrequest by the user either accepting or declining the request.
func (r *FriendRequestRepository) HandleFriendRequest(requesterID, receiverID int64) (*models.FriendRequest, error) {
	var friendRequest models.FriendRequest
	err := r.db.Where("requester_id = ? AND receiver_id = ?", requesterID, receiverID).First(&friendRequest).Error
	if err != nil {
		return nil, err
	}
	return &friendRequest, nil
}

// Notify sends a notification to the user.
func (r *FriendRequestRepository) Notify(userID int64, message string) error {
	// Send a notification to the user
	return nil
}

// Delete deletes a friend request.
func (r *FriendRequestRepository) Delete(fr *models.FriendRequest) error {
	return r.db.Delete(fr).Error
}

// CreateFriendshipEntry creates a new friendship entry.
func (r *FriendRequestRepository) CreateFriendshipEntry(userID, friendUserID int64) error {
	fe := models.FriendshipEntry{
		UserID:       int(userID),
		FriendUserID: int(friendUserID),
	}
	return r.db.Create(&fe).Error
}

// Update updates the friend request status.
func (r *FriendRequestRepository) Update(fr *models.FriendRequest) error {
	return r.db.Save(fr).Error
}
