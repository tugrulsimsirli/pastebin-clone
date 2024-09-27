package db

import (
	"log"
	"pastebin-clone/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=pastebin password=pastebin dbname=pastebin port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	log.Println("Database connection established")
}

// Migrasyon işlemi
func MigrateDB() {
	DB.AutoMigrate(&models.User{}, &models.Snippet{})
	log.Println("Database migration completed")
}
