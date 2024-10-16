package budgets

import (
	"github.com/dot-slash-ann/home-services-api/categories"
	"github.com/dot-slash-ann/home-services-api/entities"
	"gorm.io/gorm"
)

type BudgetsService interface {
	Create(CreateBudgetDto) (entities.Budget, error)
	FindAll() ([]entities.Budget, error)
	FindOne(string) (entities.Budget, error)
	FindByName(string) (entities.Budget, error)
	Delete(string) (entities.Budget, error)
	AddCategory(AddCategoryDto, string) (entities.Budget, error)
}

type BudgetsServiceImpl struct {
	database          *gorm.DB
	categoriesService categories.CategoriesService
}

func NewBudgetsService(database *gorm.DB, categoriesService categories.CategoriesService) *BudgetsServiceImpl {
	return &BudgetsServiceImpl{
		database:          database,
		categoriesService: categoriesService,
	}
}

func (service *BudgetsServiceImpl) Create(createBudgetDto CreateBudgetDto) (entities.Budget, error) {
	budget := entities.Budget{
		Name: createBudgetDto.Name,
	}

	if result := service.database.Create(&budget); result.Error != nil {
		return entities.Budget{}, result.Error
	}

	return budget, nil
}

func (service *BudgetsServiceImpl) FindAll() ([]entities.Budget, error) {
	var budgetsList []entities.Budget

	if results := service.database.Preload("Categories").Find(&budgetsList); results.Error != nil {
		return []entities.Budget{}, results.Error
	}

	return budgetsList, nil
}

func (service *BudgetsServiceImpl) FindOne(id string) (entities.Budget, error) {
	var budget entities.Budget

	if results := service.database.Preload("Categories").First(&budget, id); results.Error != nil {
		return entities.Budget{}, results.Error
	}

	return budget, nil
}

func (service *BudgetsServiceImpl) FindByName(name string) (entities.Budget, error) {
	var budget entities.Budget

	if results := service.database.Preload("Categories").First(&budget, "name = ?", name); results.Error != nil {
		return entities.Budget{}, results.Error
	}

	return budget, nil
}

func (service *BudgetsServiceImpl) Delete(id string) (entities.Budget, error) {
	budget, err := service.FindOne(id)

	if err != nil {
		return entities.Budget{}, err
	}

	if results := service.database.Delete(&entities.Budget{}, id); results.Error != nil {
		return entities.Budget{}, results.Error
	}

	return budget, nil
}

func (service *BudgetsServiceImpl) AddCategory(addCategoryDto AddCategoryDto, id string) (entities.Budget, error) {
	var budget entities.Budget

	if results := service.database.Preload("Categories").First(&budget, id); results.Error != nil {
		return entities.Budget{}, results.Error
	}

	category, err := service.categoriesService.FindByName(addCategoryDto.CategoryName)

	if err != nil {
		return entities.Budget{}, err
	}

	if err := service.database.Model(&budget).Association("Categories").Append(&category); err != nil {
		return entities.Budget{}, err
	}

	return budget, nil
}
