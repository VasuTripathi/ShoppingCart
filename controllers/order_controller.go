package controllers

import (
	"net/http"
	"shopping-cart-backend/database"
	"shopping-cart-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateOrder converts a cart into an order
func CreateOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input struct {
		CartID uint `json:"cart_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify the cart belongs to the user and is active
	var cart models.Cart
	if err := database.DB.Where("id = ? AND user_id = ? AND status = ?", input.CartID, userID, "active").First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Active cart not found or does not belong to user"})
		return
	}

	// Safely assert the type of userID before creating the order
	uid, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type in token"})
		return
	}

	// Create the order
	order := models.Order{UserID: uid, CartID: cart.ID}
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Update the cart status to "converted"
	database.DB.Model(&cart).Update("status", "converted")

	// Reload order with cart details for the response
	database.DB.Preload("Cart.CartItems.Item").First(&order, order.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully!", "order": order})
}

// GetOrders lists all orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("Cart.CartItems.Item").Preload("User").Find(&orders)
	c.JSON(http.StatusOK, orders)
}
