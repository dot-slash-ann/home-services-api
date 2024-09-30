package TagsController

import (
	"net/http"

	TagsDto "github.com/dot-slash-ann/home-services-api/dtos/tags"
	TagsService "github.com/dot-slash-ann/home-services-api/services/tags"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var createTagDto TagsDto.CreateTagDto

	if err := c.ShouldBind(&createTagDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": "expected request body",
			"code":    400,
		})

		return
	}

	tag, err := TagsService.Create(createTagDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":   tag.ID,
			"name": tag.Name,
		},
	})
}

func FindAll(c *gin.Context) {
	tags, err := TagsService.FindAll()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	results := make([]gin.H, 0, len(tags))

	for _, tag := range tags {
		results = append(results, gin.H{
			"id":   tag.ID,
			"name": tag.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func FindOne(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	tag, err := TagsService.FindOne(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":   tag.ID,
			"name": tag.Name,
		},
	})
}

func Update(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	var updateTagDto TagsDto.UpdateTagDto

	if err := c.ShouldBind(&updateTagDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": "expected request body",
			"code":    400,
		})

		return
	}

	tag, err := TagsService.Update(id, updateTagDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":   tag.ID,
			"name": tag.Name,
		},
	})
}

func Delete(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	tag, err := TagsService.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": gin.H{
			"id":   tag.ID,
			"name": tag.Name,
		},
	})
}
