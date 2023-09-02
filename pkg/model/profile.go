package model

import "gorm.io/gorm"

// Profile for user profile
type Profile struct {
	UserID           string `json:"userId" gorm:"key"`
	CurrentRequestID string `json:"currentRequestId"`
	FileName         string `json:"fileName"`
	IsDarkMode       bool   `json:"isDarkMode"`
	gorm.Model
}
