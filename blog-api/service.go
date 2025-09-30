package main

import (
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
	if err := database.DB.Find("id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func CreateBlog(blog *models.Blog) error {
	if err := database.DB.Create(blog).Error; err != nil {
		return err
	}

	return nil
}
