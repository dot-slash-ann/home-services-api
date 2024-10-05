package tags

import (
	"errors"
	"net/http"

	tagsDto "github.com/dot-slash-ann/home-services-api/dtos/tags"
	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/dot-slash-ann/home-services-api/services/tags"
	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService tags.TagsService
}

func NewTagsController(service tags.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

func (controller *TagsController) Create(c *gin.Context) {
	var createTagDto tagsDto.CreateTagDto

	if err := c.ShouldBind(&createTagDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	tag, err := controller.tagsService.Create(createTagDto)

	if err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": tagsDto.TagToJson(tag),
	})
}

func (controller *TagsController) FindAll(c *gin.Context) {
	tags, err := controller.tagsService.FindAll()

	if err != nil {
		httpErr := httpErrors.InternalServerError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tagsDto.ManyTagsToJson(tags),
	})
}

func (controller *TagsController) FindOne(c *gin.Context) {
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

	tag, err := controller.tagsService.FindOne(id)

	if err != nil {
		httpErr := httpErrors.NotFoundError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tagsDto.TagToJson(tag),
	})
}

func (controller *TagsController) Update(c *gin.Context) {
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
	var updateTagDto tagsDto.UpdateTagDto

	if err := c.ShouldBind(&updateTagDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	tag, err := controller.tagsService.Update(id, updateTagDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": tagsDto.TagToJson(tag),
	})
}

func (controller *TagsController) Delete(c *gin.Context) {
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

	tag, err := controller.tagsService.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tagsDto.TagToJson(tag),
	})
}
