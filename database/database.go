package database

import (
	"log"
	"shopping-cart-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection and runs migrations
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("shopping_cart.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	log.Println("Running Migrations")
	err = database.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.CartItem{}, &models.Order{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = database
}
