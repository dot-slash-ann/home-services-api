package vendors

import (
	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/gin-gonic/gin"
)

func VendorToJson(vendor entities.Vendor) gin.H {
	return gin.H{
		"id":   vendor.ID,
		"name": vendor.Name,
	}
}

func ManyVendorsToJson(vendors []entities.Vendor) []gin.H {
	results := make([]gin.H, 0, len(vendors))

	for _, vendor := range vendors {
		results = append(results, VendorToJson(vendor))
	}

	return results
}
