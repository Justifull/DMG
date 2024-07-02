package database

import (
	"DMG/structs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB // Storing the current database session for better access
)

// InitDatabase initializes the database with its tables
func InitDatabase() error {
	DB, _ = gorm.Open(sqlite.Open("./database/database.db"), &gorm.Config{})

	// Enable foreign keys
	db, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object: %v", err)
	}
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		log.Fatalf("failed to enable foreign keys: %v", err)
	}

	// Create the tables, missing foreign keys, constraints, columns and indexes
	err = DB.AutoMigrate(&structs.Client{}, &structs.Ride{}, &structs.Driver{}, &structs.Waypoint{})
	return err
}
