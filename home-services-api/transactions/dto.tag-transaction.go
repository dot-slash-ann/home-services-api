package transactions

type TagTransactionDto struct {
	TagName string `json:"tag_name" binding:"required"`
}
