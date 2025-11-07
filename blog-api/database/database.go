package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Rioba-Ian/blog-api/env"
	"github.com/Rioba-Ian/blog-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println("dbHost", env.PostgresDefaultDatabase)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to postgres container db", err)
	}

	log.Println("Database connection successfully established.")

	DB.AutoMigrate(&models.Blog{})

	if err = SeedDatabase(DB); err != nil {
		log.Fatalln("Failed to seed database", err)
	}

}

func SeedDatabase(db *gorm.DB) error {
	blogs := []models.Blog{
		{
			Title:    "My First Blog Post",
			Content:  "This is the content of my first blog post.",
			Category: "tech",
		},
		{Title: "My Second Blog Post", Content: "This is the content of my second blog post.", Category: "tech"},
	}

	for _, blog := range blogs {
		err := db.Create(&blog).Error
		if err != nil {
			return err
		}
	}

	return nil
}
