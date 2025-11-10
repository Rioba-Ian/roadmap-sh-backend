package main

import (
	"fmt"
	"log"
	"net/http"

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
// @Success 200 [array] Returns  array of all posts
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router / [get]

func GetPosts(ctx *gin.Context) {
	var blogs []models.Blog

	if err := GetAllBlogs(&blogs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

//	@swagger:route	GET /posts/{id} posts getPost
// Get a post by ID
// responses:
//   200: postResponse
//   400: errorResponse

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

//	@swagger:route	POST /posts posts createPost
// Create a new post
// responses:
//   201: postResponse
//   400: errorResponse

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
