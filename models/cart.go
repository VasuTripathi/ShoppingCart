package models

import "gorm.io/gorm"

// Cart holds items for a user before an order is placed
type Cart struct {
	gorm.Model
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"foreignKey:UserID" json:"-"`
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"items"`
	Status    string     `json:"status"` // e.g., "active", "converted"
}

// CartItem links an Item to a Cart
type CartItem struct {
	gorm.Model
	CartID uint `json:"cart_id"`
	ItemID uint `json:"item_id"`
	Item   Item `gorm:"foreignKey:ItemID" json:"item"`
}
