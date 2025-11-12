package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rioba-Ian/blog-api/httputil"
	"github.com/Rioba-Ian/blog-api/models"
	"github.com/gin-gonic/gin"
)

// GetPosts godoc
// @Summary Get All Posts
// @Schemes
// @Description Get all posts and details
// @Tags home
// @Accept json
// @Produce json
// @Success 200 {object}  []models.Blog
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /posts/ [get]
func GetPosts(ctx *gin.Context) {
	var blogs []models.Blog

	if err := GetAllBlogs(&blogs); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

// GetPost godoc
// @Summary Get a Post by ID
// @Schemes
// @Description Get a post by ID
// @Tags home
// @Accept json
// @Produce json
// @Success 200 {object}  models.Blog
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /posts/{id} [get]
func GetPost(ctx *gin.Context) {
	var blog models.Blog
	id := ctx.Param("id")

	if err := GetBlogByID(&blog, id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Blog Post not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

// CreatePost godoc
// @Summary Create a new post
// @Schemes
// @Description Create a new post
// @Tags home
// @Accept json
// @Produce json
// @Success 201 {object}  models.Blog
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /posts/ [post]
func CreatePost(ctx *gin.Context) {
	var blog models.Blog

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Could not create post: ", err)})
		return
	}

	if err := CreateBlog(&blog); err != nil {
		log.Fatalln("could not save post to db", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, blog)
}

// DeletePost godoc
// @Summary Delete a post
// @Schemes
// @Description Delete a post
// @Tags home
// @Accept json
// @Produce json
// @Success 200 {string}  string
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /posts/{id} [delete]
func DeletePost(ctx *gin.Context) {
	var blog models.Blog
	id := ctx.Param("id")

	if err := Delete(&blog, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server error occured",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// UpdatePost godoc
// @Summary Update a post
// @Schemes
// @Description Update a post
// @Tags home
// @Accept json
// @Produce json
// @Success 200 {string}  string
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /posts/{id} [put]
func UpdatePost(ctx *gin.Context) {
	var blog models.Blog
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Could not find post: ", err)})
		return
	}

	fmt.Println("The blog post for update::", blog)

	if err := Update(&blog, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server error occured",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
