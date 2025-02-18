package services

import (
	"app/internal/dal"
	"app/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) GetUserByID(userID int64) (*models.User, error) {
	return dal.GetUserByID(s.db, userID)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return dal.GetUserByUsername(s.db, username)
}
