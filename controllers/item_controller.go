package controllers

import (
	"net/http"
	"shopping-cart-backend/database"
	"shopping-cart-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateItem adds a new item to the store
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&item)
	c.JSON(http.StatusOK, item)
}

// GetItems lists all available items
func GetItems(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}
