package budgets

import (
	"errors"
	"net/http"

	budgetsDto "github.com/dot-slash-ann/home-services-api/dtos/budgets"
	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/dot-slash-ann/home-services-api/services/budgets"
	"github.com/gin-gonic/gin"
)

type BudgetsController struct {
	budgetsService budgets.BudgetsService
}

func NewBudgetsController(service budgets.BudgetsService) *BudgetsController {
	return &BudgetsController{
		budgetsService: service,
	}
}

func (controller *BudgetsController) Create(c *gin.Context) {
	var createBudgetDto budgetsDto.CreateBudgetDto

	if err := c.ShouldBind(&createBudgetDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	budget, err := controller.budgetsService.Create(createBudgetDto)

	if err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": budgetsDto.BudgetToJson(budget),
	})
}

func (controller *BudgetsController) FindAll(c *gin.Context) {
	budgets, err := controller.budgetsService.FindAll()

	if err != nil {
		httpErr := httpErrors.InternalServerError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": budgetsDto.ManyBudgetsToJson(budgets),
	})
}

func (controller *BudgetsController) FindOne(c *gin.Context) {
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

	budget, err := controller.budgetsService.FindOne(id)

	if err != nil {
		httpErr := httpErrors.NotFoundError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": budgetsDto.BudgetToJson(budget),
	})
}

func (controller *BudgetsController) Delete(c *gin.Context) {
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

	budget, err := controller.budgetsService.Delete(id)

	if err != nil {
		httpErr := httpErrors.NotFoundError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": budgetsDto.BudgetToJson(budget),
	})
}

func (controller *BudgetsController) AddCategory(c *gin.Context) {
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

	var addCategoryDto budgetsDto.AddCategoryDto

	if err := c.ShouldBind(&addCategoryDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	budget, err := controller.budgetsService.AddCategory(addCategoryDto, id)

	if err != nil {
		httpErr := httpErrors.InternalServerError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": budgetsDto.BudgetToJson(budget),
	})
}
