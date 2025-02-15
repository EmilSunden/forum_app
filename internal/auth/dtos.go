package auth

import "github.com/jinzhu/gorm"

type AuthRequest struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
