package vendors

type CreateVendorDto struct {
	Name string `json:"name" binding:"required"`
}
