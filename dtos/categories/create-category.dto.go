package CategoriesDto

type CreateCategoryDto struct {
	Name string `json:"name" binding:"required"`
}
