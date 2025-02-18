package models

import "gorm.io/gorm"

// RequestLog represents a log entry for an incoming HTTP request.
type RequestLog struct {
	gorm.Model
	Method     string `gorm:"size:10"`
	URL        string `gorm:"size:255"`
	StatusCode int
	ElapsedMS  int64
	Body       string `gorm:"type:text"`
}
