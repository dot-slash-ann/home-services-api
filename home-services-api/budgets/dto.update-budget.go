package budgets

type UpdateBudgetDto struct {
	Name string `json:"name" binding:"required"`
}
