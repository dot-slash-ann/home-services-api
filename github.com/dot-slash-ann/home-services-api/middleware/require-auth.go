package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	UsersService "github.com/dot-slash-ann/home-services-api/services/users"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	// TODO: this is a bad status code
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		user, err := UsersService.FindOne(fmt.Sprint(claims["sub"].(float64)))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
