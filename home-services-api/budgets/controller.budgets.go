package budgets

import (
	"errors"
	"net/http"

	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/gin-gonic/gin"
)

type BudgetsController struct {
	budgetsService BudgetsService
}

func NewBudgetsController(service BudgetsService) *BudgetsController {
	return &BudgetsController{
		budgetsService: service,
	}
}

func (controller *BudgetsController) Create(c *gin.Context) {
	var createBudgetDto CreateBudgetDto

	if err := c.ShouldBind(&createBudgetDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	if budget, err := controller.budgetsService.FindByName(createBudgetDto.Name); err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"data": BudgetToJson(budget),
		})
		return
	}

	budget, err := controller.budgetsService.Create(createBudgetDto)

	if err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": BudgetToJson(budget),
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
		"data": ManyBudgetsToJson(budgets),
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
		"data": BudgetToJson(budget),
	})
}

func (controller *BudgetsController) Update(c *gin.Context) {
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

	var updateBudgetDto UpdateBudgetDto

	if err := c.ShouldBind(&updateBudgetDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	budget, err := controller.budgetsService.Update(id, updateBudgetDto)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": BudgetToJson(budget),
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
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": BudgetToJson(budget),
	})
}
