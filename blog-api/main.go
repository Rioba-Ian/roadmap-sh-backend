package main

import (
	"net/http"

	"github.com/Rioba-Ian/blog-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to blog api service.",
		})
	})

	router.GET("/posts", GetPosts)
	router.GET("/posts/:id", GetPost)
	router.POST("/posts/", CreatePost)

	router.Run(":8080")

}
