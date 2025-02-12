package services

import (
	"app/internal/models"
	"errors"

	"gorm.io/gorm"
)

func UserExists(db *gorm.DB, username string) (bool, error) {
	var user models.User

	// Try to find the first record with the matching username.
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, so user does not exist
			return false, nil
		}
		return false, err
	}

	// Record found, so user exists
	return true, nil
}
