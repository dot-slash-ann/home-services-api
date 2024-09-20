package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dot-slash-ann/home-services-api/dtos"
	"github.com/dot-slash-ann/home-services-api/services"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var createTransactionDto dtos.CreateTransactionDto

	if err := json.NewDecoder(c.Request.Body).Decode(&createTransactionDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	transaction, error := services.TransactionsCreate(createTransactionDto)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error,
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
	transactions, error := services.TransactionsFindAll()

	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	results := []gin.H{}

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

	transaction, error := services.TransactionsFindOne(id)

	log.Default().Println("transaction: ", transaction)
	log.Default().Println("error: ", error)

	if error != nil {
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

	var updateTransactionDto dtos.UpdateTransactionDto

	if err := json.NewDecoder(c.Request.Body).Decode(&updateTransactionDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	transaction, error := services.TransactionsUpdate(id, updateTransactionDto)

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

	transaction, error := services.TransactionsDelete(id)

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
