package categories_test

import (
	"fmt"
	"log"
	"testing"

	categoriesDto "github.com/dot-slash-ann/home-services-api/dtos/categories"
	"github.com/dot-slash-ann/home-services-api/entities/categories"
	"github.com/dot-slash-ann/home-services-api/entities/tags"
	"github.com/dot-slash-ann/home-services-api/entities/transactions"
	"github.com/dot-slash-ann/home-services-api/entities/users"
	categoriesService "github.com/dot-slash-ann/home-services-api/services/categories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to in-memory SQLite database: %v", err)
	}

	db.AutoMigrate(&categories.Category{})
	db.AutoMigrate(&tags.Tag{})
	db.AutoMigrate(&transactions.Transaction{})
	db.AutoMigrate(&users.User{})

	return db
}

func getService() (categoriesService.CategoriesService, *gorm.DB) {
	db := setupTestDB()
	categoriesService := categoriesService.NewCategoriesService(db)

	return categoriesService, db
}

func create(t *testing.T, service categoriesService.CategoriesService, names []string) []categories.Category {
	categoriesList := make([]categories.Category, len(names))

	for _, name := range names {
		createCategoryDto := categoriesDto.CreateCategoryDto{
			Name: name,
		}

		category, err := service.Create(createCategoryDto)

		if err != nil {
			t.Errorf("failed creating categories: %v", err)
		}

		categoriesList = append(categoriesList, category)
	}

	return categoriesList
}

func TestCategoriesServiceCreate(t *testing.T) {
	categoriesService, db := getService()

	createCategoryDto := categoriesDto.CreateCategoryDto{
		Name: "mock category",
	}

	category, err := categoriesService.Create(createCategoryDto)

	var count int64
	db.Model(&categories.Category{}).Count(&count)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, "mock category", category.Name)
	assert.Equal(t, int64(1), count)
}

func TestCategoriesServiceFindAll(t *testing.T) {
	categoriesService, _ := getService()

	names := make([]string, 0, 5)

	for i := 0; i < 5; i++ {
		names = append(names, fmt.Sprintf("mock category %v", i))
	}

	create(t, categoriesService, names)

	categories, err := categoriesService.FindAll()

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, 5, len(categories))

	for i, category := range categories {
		assert.Equal(t, fmt.Sprintf("mock category %v", i), category.Name)
	}
}

func TestCategoriesServiceFindOne(t *testing.T) {
	categoriesService, _ := getService()

	create(t, categoriesService, []string{"mock category"})

	category, err := categoriesService.FindOne("1")

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, uint(1), category.ID)
	assert.Equal(t, "mock category", category.Name)
}

func TestCategoriesServiceFindByName(t *testing.T) {
	categoriesService, _ := getService()

	create(t, categoriesService, []string{"get by name"})

	category, err := categoriesService.FindByName("get by name")

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, uint(1), category.ID)
	assert.Equal(t, "get by name", category.Name)
}

func TestCategoriesServiceUpdate(t *testing.T) {
	categoriesService, _ := getService()

	create(t, categoriesService, []string{"old name"})

	category, _ := categoriesService.FindOne("1")

	assert.Equal(t, "old name", category.Name)

	category, err := categoriesService.Update("1", categoriesDto.UpdateCategoryDto{
		Name: "new name",
	})

	assert.Nil(t, err)

	category, _ = categoriesService.FindOne("1")

	assert.Equal(t, "new name", category.Name)
}

func TestCategoriesServiceDelete(t *testing.T) {
	categoriesService, _ := getService()

	create(t, categoriesService, []string{"to be deleted"})

	category, err := categoriesService.Delete("1")

	assert.Nil(t, err)
	assert.Equal(t, "to be deleted", category.Name)

	category, err = categoriesService.FindOne("1")

	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
}
