package categories

import (
	"errors"
	"net/http"

	categoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
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

	if err := c.ShouldBind(&createCategoryDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	category, err := controller.categoriesService.Create(createCategoryDto)

	if err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": categoriesDto.CategoryToJson(category),
	})
}

func (controller *CategoriesController) FindAll(c *gin.Context) {
	categories, err := controller.categoriesService.FindAll()

	if err != nil {
		httpErr := httpErrors.InternalServerError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesDto.ManyCategoriesToJson(categories),
	})
}

func (controller *CategoriesController) FindOne(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		httpErr := httpErrors.BadRequestError(errors.New("id arg not found"), nil)

		c.Error(httpErr)

		return
	}

	if !lib.IsNumeric(id) {
		httpErr := httpErrors.BadRequestError(errors.New("id must be an integer"), nil)

		c.Error(httpErr)

		return
	}

	category, err := controller.categoriesService.FindOne(id)

	if err != nil {
		httpErr := httpErrors.NotFoundError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesDto.CategoryToJson(category),
	})
}

func (controller *CategoriesController) Update(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		httpErr := httpErrors.BadRequestError(errors.New("id arg not found"), nil)

		c.Error(httpErr)

		return
	}

	if !lib.IsNumeric(id) {
		httpErr := httpErrors.BadRequestError(errors.New("id must be an integer"), nil)

		c.Error(httpErr)

		return
	}

	var updateCategoryDto categoriesDto.UpdateCategoryDto

	if err := c.ShouldBind(&updateCategoryDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

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
	id, found := c.Params.Get("id")

	if !found {
		httpErr := httpErrors.BadRequestError(errors.New("id arg not found"), nil)

		c.Error(httpErr)

		return
	}

	if !lib.IsNumeric(id) {
		httpErr := httpErrors.BadRequestError(errors.New("id must be an integer"), nil)

		c.Error(httpErr)

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
