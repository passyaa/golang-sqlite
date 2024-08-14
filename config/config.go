package config

import (
	"golangApp/models"
	"log"

	"golang.org/x/crypto/bcrypt"
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

	// Seed initial data
	SeedData()
}

func SeedData() {
	// Check if there are any groups, if not, insert initial data
	var count int64
	DB.Model(&models.Group{}).Count(&count)
	if count == 0 {
		// Insert initial groups
		adminGroup := models.Group{Name: "Admin", Description: "Administrator group"}
		userGroup := models.Group{Name: "User", Description: "Regular user group"}
		DB.Create(&adminGroup)
		DB.Create(&userGroup)
		log.Println("Seeded initial groups")

		// Insert initial users and assign groups
		users := []models.User{
			{Username: "admin", Email: "admin@example.com", Password: hashPassword("admin"), FirstName: "Admin", LastName: "User", Groups: []models.Group{adminGroup}},
			{Username: "user", Email: "user@example.com", Password: hashPassword("user"), FirstName: "Regular", LastName: "User", Groups: []models.Group{userGroup}},
		}
		for _, user := range users {
			DB.Create(&user)
		}
		log.Println("Seeded initial users with group assignments")
	}
}

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
