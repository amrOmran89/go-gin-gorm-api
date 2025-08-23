package controllers

import (
	"go-gin-gorm/initializers"
	"go-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	// create posts variable
	var posts []models.Post

	// fetch the posts from the DB and assign it to the posts variable
	initializers.DB.Find(&posts)

	// send the posts to the api
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func CreatePost(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
	}

	c.JSON(http.StatusCreated, gin.H{"created": true})
}

func GetPostById(c *gin.Context) {

	var id = c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{"post": post})
}
