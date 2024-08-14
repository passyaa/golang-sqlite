package config

import (
	"golangApp/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Tentukan path untuk database SQLite
	dbPath := "golangApp.db"

	// Buka koneksi ke database SQLite
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQLite database:", err)
	}

	// Auto migrate tables
	DB.AutoMigrate(&models.User{}, &models.Group{})

	log.Println("Connected to SQLite database successfully")
}
