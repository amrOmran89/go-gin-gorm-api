package controllers

import (
	models "go-gin-gorm/entities"
	"go-gin-gorm/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	// create posts variable
	var posts []models.PostEntity

	// fetch the posts from the DB and assign it to the posts variable
	initializers.DB.Find(&posts)

	// send the posts to the api
	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.PostEntity{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
	}

	c.JSON(http.StatusCreated, gin.H{"post": "created"})
}

func GetPostById(c *gin.Context) {

	var id = c.Param("id")
	var post models.PostEntity
	initializers.DB.First(&post, id)

	c.JSON(http.StatusOK, post)
}
