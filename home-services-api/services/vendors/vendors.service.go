package vendors

import (
	vendorsDto "github.com/dot-slash-ann/home-services-api/dtos/vendors"
	"github.com/dot-slash-ann/home-services-api/entities"
	"gorm.io/gorm"
)

type VendorsService interface {
	Create(vendorsDto.CreateVendorDto) (entities.Vendor, error)
	FindAll() ([]entities.Vendor, error)
	FindOne(string) (entities.Vendor, error)
	FindByName(string) (entities.Vendor, error)
	Update(string, vendorsDto.UpdateVendorDto) (entities.Vendor, error)
	Delete(string) (entities.Vendor, error)
}

type VendorsServiceImpl struct {
	database *gorm.DB
}

func NewVendorsService(database *gorm.DB) *VendorsServiceImpl {
	return &VendorsServiceImpl{
		database: database,
	}
}

func (service *VendorsServiceImpl) Create(createVendorDto vendorsDto.CreateVendorDto) (entities.Vendor, error) {
	vendor := entities.Vendor{
		Name: createVendorDto.Name,
	}

	if result := service.database.Create(&vendor); result.Error != nil {
		return entities.Vendor{}, result.Error
	}

	return vendor, nil
}

func (service *VendorsServiceImpl) FindAll() ([]entities.Vendor, error) {
	var vendorsList []entities.Vendor

	if result := service.database.Find(&vendorsList); result.Error != nil {
		return []entities.Vendor{}, result.Error
	}

	return vendorsList, nil
}

func (service *VendorsServiceImpl) FindOne(id string) (entities.Vendor, error) {
	var vendor entities.Vendor

	if result := service.database.First(&vendor, id); result.Error != nil {
		return entities.Vendor{}, result.Error
	}

	return vendor, nil
}

func (service *VendorsServiceImpl) FindByName(name string) (entities.Vendor, error) {
	var vendor entities.Vendor

	if result := service.database.First(&vendor, "name = ?", name); result.Error != nil {
		return entities.Vendor{}, result.Error
	}

	return vendor, nil
}

func (service *VendorsServiceImpl) Update(id string, updateVendorDto vendorsDto.UpdateVendorDto) (entities.Vendor, error) {
	var vendor entities.Vendor

	updatedVendor := entities.Vendor{
		Name: updateVendorDto.Name,
	}

	if result := service.database.First(&vendor, id); result.Error != nil {
		return entities.Vendor{}, result.Error
	}

	if result := service.database.Model(&vendor).Updates(updatedVendor); result.Error != nil {
		return entities.Vendor{}, result.Error
	}

	return vendor, nil
}

func (service *VendorsServiceImpl) Delete(id string) (entities.Vendor, error) {
	var vendor entities.Vendor

	if result := service.database.First(&vendor, id); result.Error != nil {
		return entities.Vendor{}, result.Error
	}

	if result := service.database.Delete(&entities.Vendor{}, id); result.Error != nil {
		return entities.Vendor{}, result.Error
	}

	return vendor, nil
}
