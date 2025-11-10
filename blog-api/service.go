package main

import (
	"fmt"
	"log"

	"github.com/Rioba-Ian/blog-api/database"

	"github.com/Rioba-Ian/blog-api/models"
)

func GetAllBlogs(blogs *[]models.Blog) error {
	if err := database.DB.Find(blogs).Error; err != nil {
		return err
	}

	return nil
}

func GetBlogByID(blog *models.Blog, id string) error {
	if err := database.DB.Where("id = ?", id).First(&blog).Error; err != nil {
		fmt.Println("Error getting blog by id", id, err)
		fmt.Errorf("Record not found", blog, id)
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
