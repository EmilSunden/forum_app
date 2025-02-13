package services

import (
	"app/internal/models"
	"errors"

	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, userID int64) (*models.User, error) {
	user := models.User{}

	// Try to find the first record with the matching username
	err := db.Where("id= ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, so user does not exist
			return nil, nil
		}
		return nil, err
	}
	// Record found, so user exists
	return &user, nil
}

func GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	user := models.User{}

	// Try to find the first record with the matching username
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, so user does not exist
			return nil, nil
		}
		return nil, err
	}
	// Record found, so user exists
	return &user, nil
}
