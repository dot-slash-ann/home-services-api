package categories

import (
	categoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	"github.com/dot-slash-ann/home-services-api/entities/categories"
	"gorm.io/gorm"
)

type CategoriesService interface {
	Create(categoriesDto.CreateCategoryDto) (categories.Category, error)
	FindAll() ([]categories.Category, error)
	FindOne(string) (categories.Category, error)
	Update(string, categoriesDto.UpdateCategoryDto) (categories.Category, error)
	Delete(string) (categories.Category, error)
}

type CategoriesServiceImpl struct {
	database *gorm.DB
}

func NewCategoriesService(database *gorm.DB) *CategoriesServiceImpl {
	return &CategoriesServiceImpl{
		database: database,
	}
}

func (service *CategoriesServiceImpl) Create(createCategoryDto categoriesDto.CreateCategoryDto) (categories.Category, error) {
	category := categories.Category{
		Name: createCategoryDto.Name,
	}

	if result := service.database.Create(&category); result.Error != nil {
		return categories.Category{}, result.Error
	}

	return category, nil
}

func (service *CategoriesServiceImpl) FindAll() ([]categories.Category, error) {
	var categoriesList []categories.Category

	if result := service.database.Find(&categoriesList); result.Error != nil {
		return []categories.Category{}, result.Error
	}

	return categoriesList, nil
}

func (service *CategoriesServiceImpl) FindOne(id string) (categories.Category, error) {
	var category categories.Category

	if result := service.database.First(&category, id); result.Error != nil {
		return categories.Category{}, result.Error
	}

	return category, nil
}

func (service *CategoriesServiceImpl) Update(id string, updateCategoryDto categoriesDto.UpdateCategoryDto) (categories.Category, error) {
	var category categories.Category

	updatedCategory := categories.Category{
		Name: updateCategoryDto.Name,
	}

	if result := service.database.First(&category, id); result.Error != nil {
		return categories.Category{}, result.Error
	}

	if result := service.database.Model(&category).Updates(updatedCategory); result.Error != nil {
		return categories.Category{}, result.Error
	}

	return category, nil
}

func (service *CategoriesServiceImpl) Delete(id string) (categories.Category, error) {
	var category categories.Category

	if result := service.database.First(&category, id); result.Error != nil {
		return categories.Category{}, result.Error
	}

	if result := service.database.Delete(&categories.Category{}, id); result.Error != nil {
		return categories.Category{}, result.Error
	}

	return category, nil
}
