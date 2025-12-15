package database

import (
	"learn-go-in-orm/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectToMySQL establishes a connection to the MySQL database
func ConnectToMySQL(dsn string) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
		return err
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("✅ Database connected successfully")
	return nil
}

// MigrateDB performs schema migration to create/update tables
func MigrateDB() error {
	if err := DB.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return err
	}

	log.Println("✅ Database migration completed successfully")
	return nil
}

// InitDB initializes database connection and performs migration
func InitDB(dsn string) error {
	if err := ConnectToMySQL(dsn); err != nil {
		return err
	}

	if err := MigrateDB(); err != nil {
		return err
	}

	return nil
}
