package users

import "gorm.io/gorm"

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) GetUserByID(userID int64) (*User, error) {
	return GetUserByID(s.db, userID)
}

func (s *UserService) GetUserByUsername(username string) (*User, error) {
	return GetUserByUsername(s.db, username)
}
