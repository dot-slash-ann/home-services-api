package categories

import (
	"github.com/dot-slash-ann/home-services-api/entities"
	"gorm.io/gorm"
)

type CategoriesService interface {
	Create(CreateCategoryDto) (entities.Category, error)
	FindAll() ([]entities.Category, error)
	FindOne(string) (entities.Category, error)
	FindByName(string) (entities.Category, error)
	Update(string, UpdateCategoryDto) (entities.Category, error)
	Delete(string) (entities.Category, error)
}

type CategoriesServiceImpl struct {
	database *gorm.DB
}

func NewCategoriesService(database *gorm.DB) *CategoriesServiceImpl {
	return &CategoriesServiceImpl{
		database: database,
	}
}

func (service *CategoriesServiceImpl) Create(createCategoryDto CreateCategoryDto) (entities.Category, error) {
	category := entities.Category{
		Name: createCategoryDto.Name,
	}

	if result := service.database.Create(&category); result.Error != nil {
		return entities.Category{}, result.Error
	}

	return category, nil
}

func (service *CategoriesServiceImpl) FindAll() ([]entities.Category, error) {
	var categoriesList []entities.Category

	if result := service.database.Find(&categoriesList); result.Error != nil {
		return []entities.Category{}, result.Error
	}

	return categoriesList, nil
}

func (service *CategoriesServiceImpl) FindOne(id string) (entities.Category, error) {
	var category entities.Category

	if result := service.database.First(&category, id); result.Error != nil {
		return entities.Category{}, result.Error
	}

	return category, nil
}

func (service *CategoriesServiceImpl) FindByName(name string) (entities.Category, error) {
	var category entities.Category

	if result := service.database.First(&category, "name = ?", name); result.Error != nil {
		return entities.Category{}, result.Error
	}

	return category, nil
}

func (service *CategoriesServiceImpl) Update(id string, updateCategoryDto UpdateCategoryDto) (entities.Category, error) {
	var category entities.Category

	updatedCategory := entities.Category{
		Name: updateCategoryDto.Name,
	}

	if result := service.database.First(&category, id); result.Error != nil {
		return entities.Category{}, result.Error
	}

	if result := service.database.Model(&category).Updates(updatedCategory); result.Error != nil {
		return entities.Category{}, result.Error
	}

	return category, nil
}

func (service *CategoriesServiceImpl) Delete(id string) (entities.Category, error) {
	var category entities.Category

	if result := service.database.First(&category, id); result.Error != nil {
		return entities.Category{}, result.Error
	}

	if result := service.database.Delete(&entities.Category{}, id); result.Error != nil {
		return entities.Category{}, result.Error
	}

	return category, nil
}
