package TransactionsController

import (
	"encoding/json"
	"net/http"

	TransactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	TransactionsService "github.com/dot-slash-ann/home-services-api/services/transactions"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var createTransactionDto TransactionsDto.CreateTransactionDto

	if err := json.NewDecoder(c.Request.Body).Decode(&createTransactionDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": "all date fields must be valid dates in the format yyyy-mm-dd",
			"code":    400,
		})

		return
	}

	transaction, err := TransactionsService.Create(createTransactionDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":             transaction.ID,
			"transaction_on": transaction.TransactionOn.UTC().Format("2006-01-02"),
			"posted_on":      transaction.PostedOn.UTC().Format("2006-01-02"),
			"amount":         transaction.Amount,
		},
	})
}

func FindAll(c *gin.Context) {
	transactions, err := TransactionsService.FindAll()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	results := make([]gin.H, 0, len(transactions))

	for _, transaction := range transactions {
		results = append(results, gin.H{
			"id":             transaction.ID,
			"transaction_on": transaction.TransactionOn.UTC().Format("2006-01-02"),
			"posted_on":      transaction.PostedOn.UTC().Format("2006-01-02"),
			"amount":         transaction.Amount,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func FindOne(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	transaction, err := TransactionsService.FindOne(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":             transaction.ID,
			"transaction_on": transaction.TransactionOn.UTC().Format("2006-01-02"),
			"posted_on":      transaction.PostedOn.UTC().Format("2006-01-02"),
			"amount":         transaction.Amount,
		},
	})
}

func Update(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	var updateTransactionDto TransactionsDto.UpdateTransactionDto

	if err := json.NewDecoder(c.Request.Body).Decode(&updateTransactionDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	transaction, error := TransactionsService.Update(id, updateTransactionDto)

	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": gin.H{
			"id":             transaction.ID,
			"transaction_on": transaction.TransactionOn.UTC().Format("2006-01-02"),
			"posted_on":      transaction.PostedOn.UTC().Format("2006-01-02"),
			"amount":         transaction.Amount,
		},
	})
}

func Delete(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	transaction, err := TransactionsService.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": gin.H{
			"id":             transaction.ID,
			"transaction_on": transaction.TransactionOn.UTC().Format("2006-01-02"),
			"posted_on":      transaction.PostedOn.UTC().Format("2006-01-02"),
			"amount":         transaction.Amount,
		},
	})
}
