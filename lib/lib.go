package lib

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

func GetParam(c *gin.Context, param string) (string, bool) {
	value, found := c.Params.Get(param)

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	return value, found
}

func HandleDecodeTime(c *gin.Context, dto interface{}) bool {
	if err := json.NewDecoder(c.Request.Body).Decode(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": "all date fields must be in the format yyyy-mm-dd",
			"code":    400,
		})

		return false
	}

	return true
}

func HandleShouldBind(c *gin.Context, dto interface{}) bool {
	if err := c.ShouldBind(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": "expected request body",
			"code":    400,
		})

		return false
	}

	return true
}

func HandleDatabaseError(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}

	return nil
}
