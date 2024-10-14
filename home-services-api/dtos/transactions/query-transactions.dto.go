package transactions

import (
	"github.com/dot-slash-ann/home-services-api/dtos/categories"
	"github.com/dot-slash-ann/home-services-api/dtos/tags"
	"github.com/dot-slash-ann/home-services-api/dtos/vendors"
	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/gin-gonic/gin"
)

func TransactionToJson(transaction entities.Transaction) gin.H {
	return gin.H{
		"id":               transaction.ID,
		"amount":           transaction.Amount,
		"transaction_on":   transaction.TransactionOn,
		"posted_on":        transaction.PostedOn,
		"transaction_type": transaction.TransactionType,
		"category":         categories.CategoryToJson(transaction.Category),
		"tags":             tags.ManyTagsToJson(transaction.Tags),
		"vendor":           vendors.VendorToJson(transaction.Vendor),
	}
}

func ManyTransactionsToJson(transactions []entities.Transaction) []gin.H {
	results := make([]gin.H, 0, len(transactions))

	for _, transaction := range transactions {
		results = append(results, TransactionToJson(transaction))
	}

	return results
}
