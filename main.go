package main

import (
	"go-gin-gorm/controllers"
	"go-gin-gorm/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/posts", controllers.GetPost)
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts/:id", controllers.GetPostById)

	router.Run()
}
