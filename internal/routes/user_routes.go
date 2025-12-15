package routes

import (
	"learn-go-in-orm/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(apiGroup *gin.RouterGroup) {
	users := apiGroup.Group("/users")
	{
		users.GET("", handlers.GetUsers)          // List all users (with pagination)
		users.POST("", handlers.CreateUser)       // Create new user
		users.GET("/:id", handlers.GetUserByID)   // Get user by ID
		users.PUT("/:id", handlers.UpdateUser)    // Update user
		users.DELETE("/:id", handlers.DeleteUser) // Delete user
	}
}
