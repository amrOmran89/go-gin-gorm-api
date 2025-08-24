package controllers

import (
	"go-gin-gorm/initializers"
	"go-gin-gorm/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get email and password from request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read the request body"})
		return
	}

	// check for empty body or required fields
	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	passwordByte := []byte(body.Password)

	// hash the password
	hash, err := bcrypt.GenerateFromPassword(passwordByte, 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash the password"})
		return
	}

	// create the user in the database
	var user = models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create a user"})
		return
	}

	// respond
	c.JSON(http.StatusCreated, gin.H{})
}

func Login(c *gin.Context) {
	// get email and password from request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read the request body"})
		return
	}

	// check for empty body or required fields
	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	// lookup requested user
	var user models.User
	initializers.DB.Find(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	// compare sent password with saved user hash password
	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := os.Getenv("SECRET")
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"error": "failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// send the jwt token
	c.JSON(http.StatusOK, gin.H{"token": "sent in the cookies"})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "is validated "})
}
