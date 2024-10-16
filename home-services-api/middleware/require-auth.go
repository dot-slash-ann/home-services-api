package middleware

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/dot-slash-ann/home-services-api/users"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(userService users.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")

		if err != nil {
			httpErr := httpErrors.UnauthorizedError(err, nil)

			c.Error(httpErr)

			c.Abort()

			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if token == nil {
			httpErr := httpErrors.UnauthorizedError(err, nil)

			c.Error(httpErr)

			c.Abort()

			return
		}

		if err != nil {
			httpErr := httpErrors.UnauthorizedError(err, nil)

			c.Error(httpErr)

			c.Abort()

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user, err := userService.FindOne(fmt.Sprint(claims["sub"].(float64)))

			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				httpErr := httpErrors.UnauthorizedError(errors.New("login expired"), nil)

				c.Error(httpErr)

				c.Abort()

				return
			}

			if err != nil || user.SessionToken != tokenString {
				httpErr := httpErrors.ForbiddenError(nil)

				c.Error(httpErr)

				c.Abort()

				return
			}

			c.Set("user", user)
		} else {
			httpErr := httpErrors.ForbiddenError(nil)

			c.Error(httpErr)

			c.Abort()

			return
		}

		c.Next()
	}
}
