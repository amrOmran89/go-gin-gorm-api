package main

import (
	"go-gin-gorm/controllers"
	"go-gin-gorm/initializers"
	"go-gin-gorm/middleware"

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

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/users", middleware.RequireAuth, controllers.GetAllUsers)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run()
}
