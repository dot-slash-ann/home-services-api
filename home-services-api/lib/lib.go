package lib

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}

func HandleError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
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
