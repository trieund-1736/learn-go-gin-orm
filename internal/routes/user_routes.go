package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(apiGroup *gin.RouterGroup) {
	users := apiGroup.Group("/users")
	{
		users.GET("", GetUsers)          // List all users (with pagination)
		users.POST("", CreateUser)       // Create new user
		users.GET("/:id", GetUserByID)   // Get user by ID
		users.PUT("/:id", UpdateUser)    // Update user
		users.DELETE("/:id", DeleteUser) // Delete user
	}
}

// GetUsers handles GET /api/v1/users
func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "List all users"})
}

// CreateUser handles POST /api/v1/users
func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Create new user"})
}

// GetUserByID handles GET /api/v1/users/:id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get user by ID", "id": id})
}

// UpdateUser handles PUT /api/v1/users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Update user", "id": id})
}

// DeleteUser handles DELETE /api/v1/users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Delete user", "id": id})
}
