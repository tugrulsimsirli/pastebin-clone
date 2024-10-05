package db

import (
	"fmt"
	"log"
	"pastebin-clone/configs"
	models "pastebin-clone/internal/db/data-models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbConfig := configs.AppConfig.DBConfig

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	log.Println("Database connection established")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Snippet{})

	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	log.Println("Migration completed successfully")
}
