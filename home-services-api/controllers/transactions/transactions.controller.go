package TransactionsController

import (
	"net/http"

	transactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	"github.com/dot-slash-ann/home-services-api/lib"
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
		lib.HandleError(c, http.StatusBadRequest, err)

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
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsDto.ManyTransactionsToJson(transactions),
	})
}

func (controller *TransactionsController) FindOne(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	transaction, err := controller.transactionsService.FindOne(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsDto.TransactionToJson(transaction),
	})
}

func (controller *TransactionsController) Update(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
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
	id, found := lib.GetParam(c, "id")

	if !found {
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
