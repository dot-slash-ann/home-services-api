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

func TestCategoriesServiceCreate(t *testing.T) {
	db := setupTestDB()
	categoriesService := categoriesService.NewCategoriesService(db)

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
	db := setupTestDB()
	categoriesService := categoriesService.NewCategoriesService(db)

	for i := 0; i < 5; i++ {
		createCategoryDto := categoriesDto.CreateCategoryDto{
			Name: fmt.Sprintf("mock category %v", i),
		}

		_, err := categoriesService.Create(createCategoryDto)

		if err != nil {
			t.Errorf("expected no error, but got: %v", err)
		}
	}

	var count int64
	db.Model(&categories.Category{}).Count(&count)

	assert.Equal(t, int64(5), count)

	categories, err := categoriesService.FindAll()

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, 5, len(categories))

	for i, category := range categories {
		assert.Equal(t, fmt.Sprintf("mock category %v", i), category.Name)
	}
}
