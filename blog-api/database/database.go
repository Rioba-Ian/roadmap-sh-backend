package database

import (
	"log"

	"github.com/Rioba-Ian/blog-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := "host=postgres_db user=rioba password=password dbname=blog_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to postgres container db", err)
	}

	log.Println("Database connection successfully established.")

	DB.AutoMigrate(&models.Blog{})
}

func SeedDatabase(db *gorm.DB) error {
	blogs := []models.Blog{
		{Title: "My First Blog Post", Content: "This is the content of my first blog post.", Category: "tech"},
		{Title: "My Second Blog Post", Content: "This is the content of my second blog post.", Category: "tech"},
	}

	for _, blog := range blogs {
		err := db.FirstOrCreate(&blog, models.Blog{Title: blog.Title, Content: blog.Content}).Error
		if err != nil {
			return err
		}
	}

	return nil
}
