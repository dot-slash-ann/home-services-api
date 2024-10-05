package transactions

import (
	"errors"
	"net/http"

	transactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/dot-slash-ann/home-services-api/services/transactions"
	"github.com/gin-gonic/gin"
)

type TransactionsController struct {
	transactionsService transactions.TransactionsService
}

func NewTransactionsController(transactionsService transactions.TransactionsService) *TransactionsController {
	return &TransactionsController{
		transactionsService: transactionsService,
	}
}

func (controller *TransactionsController) Create(c *gin.Context) {
	var createTransactionDto transactionsDto.CreateTransactionDto

	if !lib.HandleDecodeTime(c, &createTransactionDto) {
		return
	}

	transaction, err := controller.transactionsService.Create(createTransactionDto)

	if err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error((httpErr))

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": transactionsDto.TransactionToJson(transaction),
	})
}

func (controller *TransactionsController) FindAll(c *gin.Context) {
	transactions, err := controller.transactionsService.FindAll()

	// TODO: this is a bad response code
	if err != nil {
		httpErr := httpErrors.InternalServerError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsDto.ManyTransactionsToJson(transactions),
	})
}

func (controller *TransactionsController) FindOne(c *gin.Context) {
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
	transaction, err := controller.transactionsService.FindOne(id)

	if err != nil {
		httpErr := httpErrors.NotFoundError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsDto.TransactionToJson(transaction),
	})
}

func (controller *TransactionsController) Update(c *gin.Context) {
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

	var updateTransactionDto transactionsDto.UpdateTransactionDto

	if !lib.HandleDecodeTime(c, &updateTransactionDto) {
		return
	}

	transaction, err := controller.transactionsService.Update(id, updateTransactionDto)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsDto.TransactionToJson(transaction),
	})
}

func (controller *TransactionsController) Delete(c *gin.Context) {
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

	transaction, err := controller.transactionsService.Delete(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsDto.TransactionToJson(transaction),
	})
}
