package models

import "gorm.io/gorm"

// User defines the user model
type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null" json:"username"`
	Password string `gorm:"not null"`                 // Omit password from JSON responses
	Token    string `gorm:"-" json:"token,omitempty"` // Omit token from DB, include in JSON when present
}
