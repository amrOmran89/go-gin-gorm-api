package middleware

import (
	"fmt"
	"go-gin-gorm/initializers"
	"go-gin-gorm/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode and validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		fmt.Println("two")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Check expiration date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("three")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			fmt.Println("four")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)

	} else {
		fmt.Println("five")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
