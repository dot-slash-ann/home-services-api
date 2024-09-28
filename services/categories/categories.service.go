package CategoriesService

import (
	"github.com/dot-slash-ann/home-services-api/database"
	CategoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	CategoriesEntity "github.com/dot-slash-ann/home-services-api/entities/categories"
	"github.com/dot-slash-ann/home-services-api/lib"
)

func Create(createCategoryDto CategoriesDto.CreateCategoryDto) (CategoriesEntity.Category, error) {
	category := CategoriesEntity.Category{
		Name: createCategoryDto.Name,
	}

	if err := lib.HandleDatabaseError(database.Connection.Create(&category)); err != nil {
		return CategoriesEntity.Category{}, err
	}

	return category, nil
}

func FindAll() ([]CategoriesEntity.Category, error) {
	var categories []CategoriesEntity.Category

	if err := lib.HandleDatabaseError(database.Connection.Find(&categories)); err != nil {
		return []CategoriesEntity.Category{}, err
	}

	return categories, nil
}

func FindOne(id string) (CategoriesEntity.Category, error) {
	var category CategoriesEntity.Category

	if err := lib.HandleDatabaseError(database.Connection.First(&category, id)); err != nil {
		return CategoriesEntity.Category{}, err
	}

	return category, nil
}

func Update(id string, updateCategoryDto CategoriesDto.UpdateCategoryDto) (CategoriesEntity.Category, error) {
	var category CategoriesEntity.Category

	updatedCategory := CategoriesEntity.Category{
		Name: updateCategoryDto.Name,
	}

	if err := lib.HandleDatabaseError(database.Connection.First(&category, id)); err != nil {
		return CategoriesEntity.Category{}, err
	}

	if err := lib.HandleDatabaseError(database.Connection.Model(&category).Updates(updatedCategory)); err != nil {
		return CategoriesEntity.Category{}, err
	}

	return category, nil
}

func Delete(id string) (CategoriesEntity.Category, error) {
	var category CategoriesEntity.Category

	if err := lib.HandleDatabaseError(database.Connection.First(&category, id)); err != nil {
		return CategoriesEntity.Category{}, err
	}

	if err := lib.HandleDatabaseError(database.Connection.Delete(&CategoriesEntity.Category{}, id)); err != nil {
		return CategoriesEntity.Category{}, err
	}

	return category, nil
}
