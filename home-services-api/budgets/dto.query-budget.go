package budgets

import (
	"github.com/dot-slash-ann/home-services-api/categories"
	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/gin-gonic/gin"
)

func BudgetToJson(budget entities.Budget) gin.H {
	return gin.H{
		"id":         budget.ID,
		"name":       budget.Name,
		"categories": categories.ManyCategoriesToJson(budget.Categories),
	}
}

func ManyBudgetsToJson(budgets []entities.Budget) []gin.H {
	results := make([]gin.H, 0, len(budgets))

	for _, budget := range budgets {
		results = append(results, BudgetToJson(budget))
	}

	return results
}
