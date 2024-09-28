package CategoriesDto

import (
	CategoriesEntity "github.com/dot-slash-ann/home-services-api/entities/categories"
	"github.com/gin-gonic/gin"
)

func CategoryToJson(category CategoriesEntity.Category) gin.H {
	return gin.H{
		"id":   category.ID,
		"name": category.Name,
	}
}

func ManyCategoriesToJson(categories []CategoriesEntity.Category) []gin.H {
	results := make([]gin.H, 0, len(categories))

	for _, category := range categories {
		results = append(results, CategoryToJson(category))
	}

	return results
}
