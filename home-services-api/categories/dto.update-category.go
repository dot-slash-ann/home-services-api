package categories

type UpdateCategoryDto struct {
	Name string `json:"name" binding:"required"`
}
