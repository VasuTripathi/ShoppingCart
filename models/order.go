package models

import "gorm.io/gorm"

// Order is created from a cart
type Order struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
	CartID uint `gorm:"uniqueIndex" json:"cart_id"` // Each cart can only become one order
	Cart   Cart `gorm:"foreignKey:CartID" json:"cart"`
}
