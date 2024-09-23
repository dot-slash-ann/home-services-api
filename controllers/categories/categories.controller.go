package CategoriesController

import (
	"net/http"

	CategoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	CategoriesService "github.com/dot-slash-ann/home-services-api/services/categories"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var createCategoryDto CategoriesDto.CreateCategoryDto

	if err := c.ShouldBind(&createCategoryDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": "expected request body",
			"code":    400,
		})

		return
	}

	category, err := CategoriesService.Create(createCategoryDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":   category.ID,
			"name": category.Name,
		},
	})
}

func FindAll(c *gin.Context) {
	categories, err := CategoriesService.FindAll()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	results := make([]gin.H, 0, len(categories))

	for _, category := range categories {
		results = append(results, gin.H{
			"id":   category.ID,
			"name": category.Name,
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

	category, err := CategoriesService.FindOne(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":   category.ID,
			"name": category.Name,
		},
	})
}

func Update(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	var updateCategoryDto CategoriesDto.UpdateCategoryDto

	if err := c.ShouldBind(&updateCategoryDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": "expected request body",
			"code":    400,
		})

		return
	}

	category, err := CategoriesService.Update(id, updateCategoryDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":   category.ID,
			"name": category.Name,
		},
	})
}

func Delete(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	category, err := CategoriesService.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": gin.H{
			"id":   category.ID,
			"name": category.Name,
		},
	})
}
