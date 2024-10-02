package tags

import (
	"net/http"

	tagsDto "github.com/dot-slash-ann/home-services-api/dtos/tags"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	tag, err := controller.tagsService.Create(createTagDto)

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

func (controller *TagsController) FindAll(c *gin.Context) {
	tags, err := controller.tagsService.FindAll()

	if err != nil {
		//
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

func (controller *TagsController) FindOne(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	tag, err := controller.tagsService.FindOne(id)

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

func (controller *TagsController) Update(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	var updateTagDto tagsDto.UpdateTagDto

	if err := c.ShouldBind(&updateTagDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": "expected request body",
			"code":    400,
		})

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
		"data": gin.H{
			"id":   tag.ID,
			"name": tag.Name,
		},
	})
}

func (controller *TagsController) Delete(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	tag, err := controller.tagsService.Delete(id)

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
