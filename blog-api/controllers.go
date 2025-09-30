package main

import (
	"fmt"
	"net/http"

	"github.com/Rioba-Ian/blog-api/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	var blogs []models.Blog

	if err := GetAllBlogs(&blogs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func GetPost(ctx *gin.Context) {
	var blog models.Blog
	id := ctx.Param(":id")

	if err := GetBlogByID(&blog, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func CreatePost(ctx *gin.Context) {
	var blog models.Blog

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Could not create post: ", err)})
		return
	}

	ctx.JSON(http.StatusCreated, blog)
}
