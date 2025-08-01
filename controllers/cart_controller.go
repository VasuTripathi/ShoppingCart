package controllers

import (
	"net/http"
	"shopping-cart-backend/database"
	"shopping-cart-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddToCart handles adding items to a user's cart
func AddToCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input struct {
		ItemID uint `json:"item_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user's active cart or create a new one
	var cart models.Cart
	err := database.DB.Where("user_id = ? AND status = ?", userID, "active").First(&cart).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		// --- FIX: Safe Type Assertion ---
		uid, ok := userID.(uint)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type in token"})
			return
		}
		cart = models.Cart{UserID: uid, Status: "active"}
		database.DB.Create(&cart)
		// --- END FIX ---
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error on finding cart"})
		return
	}

	// Add item to cart
	cartItem := models.CartItem{CartID: cart.ID, ItemID: input.ItemID}
	database.DB.Create(&cartItem)

	// Reload cart with items to return it
	database.DB.Preload("CartItems.Item").First(&cart, cart.ID)
	c.JSON(http.StatusOK, cart)
}

// GetCarts lists all carts (for admin purposes)
func GetCarts(c *gin.Context) {
	var carts []models.Cart
	database.DB.Preload("CartItems.Item").Preload("User").Find(&carts)
	c.JSON(http.StatusOK, carts)
}
