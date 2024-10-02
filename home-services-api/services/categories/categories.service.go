package categories

import (
	"github.com/dot-slash-ann/home-services-api/database"
	categoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	"github.com/dot-slash-ann/home-services-api/entities/categories"
	"github.com/dot-slash-ann/home-services-api/lib"
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

	if err := lib.HandleDatabaseError(database.Connection.Create(&category)); err != nil {
		return categories.Category{}, err
	}

	return category, nil
}

func (service *CategoriesServiceImpl) FindAll() ([]categories.Category, error) {
	var categoriesList []categories.Category

	if err := lib.HandleDatabaseError(database.Connection.Find(&categoriesList)); err != nil {
		return []categories.Category{}, err
	}

	return categoriesList, nil
}

func (service *CategoriesServiceImpl) FindOne(id string) (categories.Category, error) {
	var category categories.Category

	if err := lib.HandleDatabaseError(database.Connection.First(&category, id)); err != nil {
		return categories.Category{}, err
	}

	return category, nil
}

func (service *CategoriesServiceImpl) Update(id string, updateCategoryDto categoriesDto.UpdateCategoryDto) (categories.Category, error) {
	var category categories.Category

	updatedCategory := categories.Category{
		Name: updateCategoryDto.Name,
	}

	if err := lib.HandleDatabaseError(database.Connection.First(&category, id)); err != nil {
		return categories.Category{}, err
	}

	if err := lib.HandleDatabaseError(database.Connection.Model(&category).Updates(updatedCategory)); err != nil {
		return categories.Category{}, err
	}

	return category, nil
}

func (service *CategoriesServiceImpl) Delete(id string) (categories.Category, error) {
	var category categories.Category

	if err := lib.HandleDatabaseError(database.Connection.First(&category, id)); err != nil {
		return categories.Category{}, err
	}

	if err := lib.HandleDatabaseError(database.Connection.Delete(&categories.Category{}, id)); err != nil {
		return categories.Category{}, err
	}

	return category, nil
}
