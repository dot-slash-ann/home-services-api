package tags

type CreateTagDto struct {
	Name string `json:"name" binding:"required"`
}
