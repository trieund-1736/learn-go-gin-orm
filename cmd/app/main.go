package main

import (
	"learn-go-in-orm/internal/config"
	"learn-go-in-orm/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	dsn := "root:@tcp(127.0.0.1:3306)/learn_go_gin_orm?charset=utf8mb4&parseTime=True&loc=Local"
	if err := config.InitDB(dsn); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create a new Gin router
	router := gin.Default()

	// Create API apiGroup group
	apiGroup := router.Group("/api/v1")

	// Register user routes
	routes.RegisterUserRoutes(apiGroup)

	// Start the server
	router.Run(":8080")
}
