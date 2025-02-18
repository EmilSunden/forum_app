package dtos

import "gorm.io/gorm"

type AuthRequest struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
