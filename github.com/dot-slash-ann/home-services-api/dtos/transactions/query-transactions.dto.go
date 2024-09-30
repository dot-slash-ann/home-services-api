package TransactionsDto

import (
	TransactionsEntity "github.com/dot-slash-ann/home-services-api/entities/transactions"
	"github.com/gin-gonic/gin"
)

func TransactionToJson(transaction TransactionsEntity.Transaction) gin.H {
	return gin.H{
		"id":             transaction.ID,
		"transaction_on": transaction.TransactionOn,
		"posted_on":      transaction.PostedOn,
		"amount":         transaction.Amount,
		"category": gin.H{
			"id":   transaction.Category.ID,
			"name": transaction.Category.Name,
		},
	}
}

func ManyTransactionsToJson(transactions []TransactionsEntity.Transaction) []gin.H {
	results := make([]gin.H, 0, len(transactions))

	for _, transaction := range transactions {
		results = append(results, TransactionToJson(transaction))
	}

	return results
}
