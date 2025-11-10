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
		return fmt.Errorf("Record not found", blog, id)
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

func Delete(blog *models.Blog, id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&blog).Error; err != nil {
		fmt.Errorf("Record not found", blog, id)
	}

	return nil
}

func Update(blog *models.Blog, id string) error {
	if err := database.DB.Model(&blog).Updates(models.Blog{
		Base:     models.Base{},
		Title:    blog.Title,
		Content:  blog.Title,
		Category: blog.Category,
		Tags:     blog.Tags,
	}).Error; err != nil {
		log.Fatalln("Could not update post", err)
		fmt.Errorf("Could not update post", err)
	}

	return nil
}
