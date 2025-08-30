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

	postRouterV1 := router.Group("/v1/posts")
	{
		postRouterV1.GET("/", controllers.GetPost)
		postRouterV1.POST("/", controllers.CreatePost)
		postRouterV1.GET("/:id", controllers.GetPostById)
	}

	authRouterV1 := router.Group("/v1/auth")
	{
		authRouterV1.POST("/signup", controllers.SignUp)
		authRouterV1.POST("/login", controllers.Login)
	}

	userRouterV1 := router.Group("/v1")
	{
		userRouterV1.GET("/users", middleware.RequireAuth, controllers.GetAllUsers)
		userRouterV1.GET("/validate", middleware.RequireAuth, controllers.Validate)
	}

	router.Run()
}
