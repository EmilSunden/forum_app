package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"unique;not null" json:"username"`
	Password    string `gorm:"not null" json:"-"`
	DisplayName string `json:"display_name,omitempty"`
	GivenName   string `json:"given_name,omitempty"`
	FamilyName  string `json:"family_name,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
}
