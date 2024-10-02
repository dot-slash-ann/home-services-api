package CategoriesController

import (
	"net/http"

	categoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/dot-slash-ann/home-services-api/services/categories"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	categoriesService categories.CategoriesService
}

func NewCategoriesController(service categories.CategoriesService) *CategoriesController {
	return &CategoriesController{
		categoriesService: service,
	}
}

func (controller *CategoriesController) Create(c *gin.Context) {
	var createCategoryDto categoriesDto.CreateCategoryDto

	if !lib.HandleShouldBind(c, &createCategoryDto) {
		return
	}

	category, err := controller.categoriesService.Create(createCategoryDto)

	if err != nil {
		lib.HandleError(c, http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": categoriesDto.CategoryToJson(category),
	})
}

func (controller *CategoriesController) FindAll(c *gin.Context) {
	categories, err := controller.categoriesService.FindAll()

	// TODO: this is a bad response code
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesDto.ManyCategoriesToJson(categories),
	})
}

func (controller *CategoriesController) FindOne(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	category, err := controller.categoriesService.FindOne(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesDto.CategoryToJson(category),
	})
}

func (controller *CategoriesController) Update(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	var updateCategoryDto categoriesDto.UpdateCategoryDto

	if !lib.HandleShouldBind(c, &updateCategoryDto) {
		return
	}

	category, err := controller.categoriesService.Update(id, updateCategoryDto)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesDto.CategoryToJson(category),
	})
}

func (controller *CategoriesController) Delete(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	category, err := controller.categoriesService.Delete(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesDto.CategoryToJson(category),
	})
}
