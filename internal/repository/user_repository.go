package repository

import (
	"learn-go-in-orm/internal/config"
	"learn-go-in-orm/internal/models"

	"gorm.io/gorm"
)

// GetUsers retrieves all users with pagination
func GetUsers(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Count total records
	if err := config.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Fetch paginated records
	if err := config.DB.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateUser creates a new user in the database
func CreateUser(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user
func UpdateUser(id uint, updates *models.User) (*models.User, error) {
	var user models.User

	// Find user first
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	// Update user with provided data
	if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser deletes a user by ID (soft delete)
func DeleteUser(id uint) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return nil
}
