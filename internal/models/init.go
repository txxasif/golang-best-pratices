package models

import (
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection and runs migrations
func Init(database *gorm.DB) {
	DB = database

	// Run migrations
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	log.Println("Database migrated successfully!")
}
