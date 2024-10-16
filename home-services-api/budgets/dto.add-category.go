package budgets

type AddCategoryDto struct {
	CategoryName string `json:"category_name" binding:"required"`
}
