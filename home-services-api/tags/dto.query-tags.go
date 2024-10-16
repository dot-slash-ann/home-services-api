package tags

import (
	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/gin-gonic/gin"
)

func TagToJson(tag entities.Tag) gin.H {
	return gin.H{
		"id":   tag.ID,
		"name": tag.Name,
	}
}

func ManyTagsToJson(tags []entities.Tag) []gin.H {
	results := make([]gin.H, 0, len(tags))

	for _, tag := range tags {
		results = append(results, TagToJson(tag))
	}

	return results
}
