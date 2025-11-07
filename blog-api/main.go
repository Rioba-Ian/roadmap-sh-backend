package main

import (
	"log"
	"net/http"

	"github.com/Rioba-Ian/blog-api/database"
	docs "github.com/Rioba-Ian/blog-api/docs"
	"github.com/Rioba-Ian/blog-api/env"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// load env
	if err := env.Load(); err != nil {
		log.Fatal("Failed to load env variables...")
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	database.Connect()

	router.GET("/", Home)

	router.GET("/posts/", GetPosts)
	router.GET("/posts/:id", GetPost)
	router.POST("/posts/", CreatePost)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")

}

// Home godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags home
// @Accept json
// @Produce json
// @Success 200 {string} Welcome to blog api service.
// @Router / [get]
func Home(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "Welcome to blog api service.",
	})
}
