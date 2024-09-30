package CategoriesController

import (
	"net/http"

	CategoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	"github.com/dot-slash-ann/home-services-api/lib"
	CategoriesService "github.com/dot-slash-ann/home-services-api/services/categories"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var createCategoryDto CategoriesDto.CreateCategoryDto

	if !lib.HandleShouldBind(c, &createCategoryDto) {
		return
	}

	category, err := CategoriesService.Create(createCategoryDto)

	if err != nil {
		lib.HandleError(c, http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": CategoriesDto.CategoryToJson(category),
	})
}

func FindAll(c *gin.Context) {
	categories, err := CategoriesService.FindAll()

	// TODO: this is a bad response code
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": CategoriesDto.ManyCategoriesToJson(categories),
	})
}

func FindOne(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	category, err := CategoriesService.FindOne(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": CategoriesDto.CategoryToJson(category),
	})
}

func Update(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	var updateCategoryDto CategoriesDto.UpdateCategoryDto

	if !lib.HandleShouldBind(c, &updateCategoryDto) {
		return
	}

	category, err := CategoriesService.Update(id, updateCategoryDto)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": CategoriesDto.CategoryToJson(category),
	})
}

func Delete(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	category, err := CategoriesService.Delete(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": CategoriesDto.CategoryToJson(category),
	})
}
