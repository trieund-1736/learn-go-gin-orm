package services

import (
	"learn-go-in-orm/internal/models"
	"learn-go-in-orm/internal/repository"

	"gorm.io/gorm"
)

// ListUsers retrieves all users with pagination
func ListUsers(page, pageSize int) ([]models.User, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Call repository to fetch users
	users, total, err := repository.GetUsers(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateUser creates a new user
func CreateUser(user *models.User) error {
	// Call repository to create user
	if err := repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

// GetUser retrieves a user by ID
func GetUser(id uint) (*models.User, error) {
	// Validate ID
	if id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// Call repository to fetch user
	user, err := repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user
func UpdateUser(id uint, updates *models.User) (*models.User, error) {
	// Validate ID
	if id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// Call repository to update user
	user, err := repository.UpdateUser(id, updates)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id uint) error {
	// Validate ID
	if id <= 0 {
		return gorm.ErrRecordNotFound
	}

	// Call repository to delete user
	if err := repository.DeleteUser(id); err != nil {
		return err
	}

	return nil
}
