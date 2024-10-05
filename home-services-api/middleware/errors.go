package middleware

import (
	"net/http"

	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch e := err.Err.(type) {

			case httpErrors.Http:
				c.AbortWithStatusJSON(e.StatusCode, e)

			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Service Unavailable",
				})
			}
		}
	}
}
