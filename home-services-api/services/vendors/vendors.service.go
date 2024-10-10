package vendors

import (
	vendorsDto "github.com/dot-slash-ann/home-services-api/dtos/vendors"
	"github.com/dot-slash-ann/home-services-api/entities/vendors"
	"gorm.io/gorm"
)

type VendorsService interface {
	Create(vendorsDto.CreateVendorDto) (vendors.Vendor, error)
	FindAll() ([]vendors.Vendor, error)
	FindOne(string) (vendors.Vendor, error)
	FindByName(string) (vendors.Vendor, error)
	Update(string, vendorsDto.UpdateVendorDto) (vendors.Vendor, error)
	Delete(string) (vendors.Vendor, error)
}

type VendorsServiceImpl struct {
	database *gorm.DB
}

func NewVendorsService(database *gorm.DB) *VendorsServiceImpl {
	return &VendorsServiceImpl{
		database: database,
	}
}

func (service *VendorsServiceImpl) Create(createVendorDto vendorsDto.CreateVendorDto) (vendors.Vendor, error) {
	vendor := vendors.Vendor{
		Name: createVendorDto.Name,
	}

	if result := service.database.Create(&vendor); result.Error != nil {
		return vendors.Vendor{}, result.Error
	}

	return vendor, nil
}

func (service *VendorsServiceImpl) FindAll() ([]vendors.Vendor, error) {
	var vendorsList []vendors.Vendor

	if result := service.database.Find(&vendorsList); result.Error != nil {
		return []vendors.Vendor{}, result.Error
	}

	return vendorsList, nil
}

func (service *VendorsServiceImpl) FindOne(id string) (vendors.Vendor, error) {
	var vendor vendors.Vendor

	if result := service.database.First(&vendor, id); result.Error != nil {
		return vendors.Vendor{}, result.Error
	}

	return vendor, nil
}

func (service *VendorsServiceImpl) FindByName(name string) (vendors.Vendor, error) {
	var vendor vendors.Vendor

	if result := service.database.First(&vendor, "name = ?", name); result.Error != nil {
		return vendors.Vendor{}, result.Error
	}

	return vendor, nil
}

func (service *VendorsServiceImpl) Update(id string, updateVendorDto vendorsDto.UpdateVendorDto) (vendors.Vendor, error) {
	var vendor vendors.Vendor

	updatedVendor := vendors.Vendor{
		Name: updateVendorDto.Name,
	}

	if result := service.database.First(&vendor, id); result.Error != nil {
		return vendors.Vendor{}, result.Error
	}

	if result := service.database.Model(&vendor).Updates(updatedVendor); result.Error != nil {
		return vendors.Vendor{}, result.Error
	}

	return vendor, nil
}

func (service *VendorsServiceImpl) Delete(id string) (vendors.Vendor, error) {
	var vendor vendors.Vendor

	if result := service.database.First(&vendor, id); result.Error != nil {
		return vendors.Vendor{}, result.Error
	}

	if result := service.database.Delete(&vendors.Vendor{}, id); result.Error != nil {
		return vendors.Vendor{}, result.Error
	}

	return vendor, nil
}
