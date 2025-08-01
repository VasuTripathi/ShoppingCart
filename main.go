package main

import (
	"log"
	"shopping-cart-backend/controllers"
	"shopping-cart-backend/database"
	"shopping-cart-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.ConnectDatabase()

	// Initialize Router
	router := gin.Default()

	// Public Routes
	public := router.Group("/api")
	{
		public.POST("/signup", controllers.Signup)
		public.POST("/login", controllers.Login)
		public.GET("/users", controllers.GetUsers)
		public.GET("/items", controllers.GetItems)
	}

	// Authenticated Routes
	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		// In a real app, creating items might be an admin-only function
		protected.POST("/items", controllers.CreateItem)
		protected.POST("/carts", controllers.AddToCart)
		protected.GET("/carts", controllers.GetCarts)
		protected.POST("/orders", controllers.CreateOrder)
		protected.GET("/orders", controllers.GetOrders)
	}

	// Start Server
	log.Println(" Server starting on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
