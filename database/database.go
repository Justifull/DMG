package database

import (
	"DMG/structs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB // Storing the current database session for better access
)

func InitDatabase() error {
	DB, _ = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	// AutoMigrate will create the tables, missing foreign keys, constraints, columns and indexes.
	err := DB.AutoMigrate(&structs.Client{}, &structs.Ride{}, &structs.Driver{}, &structs.Waypoint{})
	return err
}
