package dal

import (
	"app/internal/models"

	"gorm.io/gorm"
)

type FriendRequestDataLayer struct {
	db *gorm.DB
}

func NewFriendRequest(db *gorm.DB) *FriendRequestDataLayer {
	return &FriendRequestDataLayer{db: db}
}

// Exists checks if a friend request already exists between two users.
func (r *FriendRequestDataLayer) Exists(requesterID, receiverID int64) (bool, error) {
	var count int64
	err := r.db.Model(&models.FriendRequest{}).Where("requester_id = ? And receiver_id = ?", requesterID, receiverID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Create creates a new friend request.
func (r *FriendRequestDataLayer) Create(fr *models.FriendRequest) error {
	return r.db.Create(fr).Error
}

// HandleFriendRequest handles a friendrequest by the user either accepting or declining the request.
func (r *FriendRequestDataLayer) HandleFriendRequest(requesterID, receiverID int64) (*models.FriendRequest, error) {
	var friendRequest models.FriendRequest
	err := r.db.Where("requester_id = ? AND receiver_id = ?", requesterID, receiverID).First(&friendRequest).Error
	if err != nil {
		return nil, err
	}
	return &friendRequest, nil
}

// Notify sends a notification to the user.
func (r *FriendRequestDataLayer) Notify(userID int64, message string) error {
	return nil
}

// Delete deletes a friend request.
func (r *FriendRequestDataLayer) Delete(fr *models.FriendRequest) error {
	return r.db.Delete(fr).Error
}

// CreateFriendshipEntry creates a new friendship entry.
func (r *FriendRequestDataLayer) CreateFriendshipEntry(userID, friendUserID int64) error {
	fe := models.FriendshipEntry{
		UserID:       int(userID),
		FriendUserID: int(friendUserID),
	}
	return r.db.Create(&fe).Error
}

// Update updates the friend request status.
func (r *FriendRequestDataLayer) Update(fr *models.FriendRequest) error {
	return r.db.Save(fr).Error
}
