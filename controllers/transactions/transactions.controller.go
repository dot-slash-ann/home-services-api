package TransactionsController

import (
	"net/http"

	TransactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	"github.com/dot-slash-ann/home-services-api/lib"
	TransactionsService "github.com/dot-slash-ann/home-services-api/services/transactions"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var createTransactionDto TransactionsDto.CreateTransactionDto

	if !lib.HandleDecodeTime(c, &createTransactionDto) {
		return
	}

	transaction, err := TransactionsService.Create(createTransactionDto)

	if err != nil {
		lib.HandleError(c, http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": TransactionsDto.TransactionToJson(transaction),
	})
}

func FindAll(c *gin.Context) {
	transactions, err := TransactionsService.FindAll()

	// TODO: this is a bad response code
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": TransactionsDto.ManyTransactionsToJson(transactions),
	})
}

func FindOne(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	transaction, err := TransactionsService.FindOne(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": TransactionsDto.TransactionToJson(transaction),
	})
}

func Update(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	var updateTransactionDto TransactionsDto.UpdateTransactionDto

	if !lib.HandleDecodeTime(c, &updateTransactionDto) {
		return
	}

	transaction, err := TransactionsService.Update(id, updateTransactionDto)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": TransactionsDto.TransactionToJson(transaction),
	})
}

func Delete(c *gin.Context) {
	id, found := lib.GetParam(c, "id")

	if !found {
		return
	}

	transaction, err := TransactionsService.Delete(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": TransactionsDto.TransactionToJson(transaction),
	})
}
