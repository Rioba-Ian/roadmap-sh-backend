package main

import (
	"fmt"
	"log"

	"github.com/Rioba-Ian/blog-api/database"
	"gorm.io/gorm"

	"github.com/Rioba-Ian/blog-api/models"
)

func GetAllBlogs(blogs *[]models.Blog) error {
	if err := database.DB.Find(blogs).Error; err != nil {
		return err
	}

	return nil
}

func GetBlogByID(blogs *[]models.Blog, id string) error {
	result := database.DB.Where("id = ?", id).Find(blogs)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Errorf("Record not found")
		} else {
			fmt.Errorf("Error retrieving user", result.Error)
		}
	}

	return nil
}

func CreateBlog(blog *models.Blog) error {
	if err := database.DB.Create(blog).Error; err != nil {
		log.Fatalln("could not create post", err)
		return err
	}

	return nil
}
