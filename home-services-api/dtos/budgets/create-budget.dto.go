package budgets

type CreateBudgetDto struct {
	Name string `json:"name" binding:"required"`
}
