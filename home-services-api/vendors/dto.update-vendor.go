package vendors

type UpdateVendorDto struct {
	Name string `json:"name" binding:"required"`
}
