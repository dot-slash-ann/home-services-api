package CategoriesService

import (
	"github.com/dot-slash-ann/home-services-api/database"
	CategoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	CategoriesEntity "github.com/dot-slash-ann/home-services-api/entities/categories"
)

func Create(createCategoryDto CategoriesDto.CreateCategoryDto) (CategoriesEntity.Category, error) {
	category := CategoriesEntity.Category{
		Name: createCategoryDto.Name,
	}

	if result := database.Connection.Create(&category); result.Error != nil {
		return CategoriesEntity.Category{}, result.Error
	}

	return category, nil
}

func FindAll() ([]CategoriesEntity.Category, error) {
	var categories []CategoriesEntity.Category

	if results := database.Connection.Find(&categories); results.Error != nil {
		return []CategoriesEntity.Category{}, results.Error
	}

	return categories, nil
}

func FindOne(id string) (CategoriesEntity.Category, error) {
	var category CategoriesEntity.Category

	if result := database.Connection.First(&category, id); result.Error != nil {
		return CategoriesEntity.Category{}, result.Error
	}

	return category, nil
}

func Update(id string, updateCategoryDto CategoriesDto.UpdateCategoryDto) (CategoriesEntity.Category, error) {
	var category CategoriesEntity.Category

	if result := database.Connection.First(&category, id); result.Error != nil {
		return CategoriesEntity.Category{}, result.Error
	}

	if result := database.Connection.Model(&category).Updates(CategoriesEntity.Category{
		Name: updateCategoryDto.Name,
	}); result.Error != nil {
		return CategoriesEntity.Category{}, result.Error
	}

	return category, nil
}

func Delete(id string) (CategoriesEntity.Category, error) {
	var category CategoriesEntity.Category

	if result := database.Connection.First(&category, id); result.Error != nil {
		return CategoriesEntity.Category{}, result.Error
	}

	database.Connection.Delete(&CategoriesEntity.Category{}, id)

	return category, nil
}
