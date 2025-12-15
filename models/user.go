package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255);not null" json:"name" binding:"required"`
	Email     string         `gorm:"column:email;type:varchar(255);uniqueIndex;not null" json:"email" binding:"required,email"`
	Password  string         `gorm:"column:password;type:varchar(255);not null" json:"password" binding:"required"`
	Phone     string         `gorm:"column:phone;type:varchar(255)" json:"phone"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}
