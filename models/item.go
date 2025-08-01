package models

import "gorm.io/gorm"

// Item defines the product model
type Item struct {
	gorm.Model
	Name   string  `json:"name"`
	Status string  `json:"status"`
	Price  float64 `json:"price"` // Added price for a more complete model
}
