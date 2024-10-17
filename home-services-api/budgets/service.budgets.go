package budgets

import (
	"github.com/dot-slash-ann/home-services-api/entities"
	"gorm.io/gorm"
)

type BudgetsService interface {
	Create(CreateBudgetDto) (entities.Budget, error)
	FindAll() ([]entities.Budget, error)
	FindOne(string) (entities.Budget, error)
	FindByName(string) (entities.Budget, error)
	Update(string, UpdateBudgetDto) (entities.Budget, error)
	Delete(string) (entities.Budget, error)
}

type BudgetsServiceImpl struct {
	database *gorm.DB
}

func NewBudgetsService(database *gorm.DB) *BudgetsServiceImpl {
	return &BudgetsServiceImpl{
		database: database,
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

	if result := service.database.Find(&budgetsList); result.Error != nil {
		return []entities.Budget{}, result.Error
	}

	return budgetsList, nil
}

func (service *BudgetsServiceImpl) FindOne(id string) (entities.Budget, error) {
	var budget entities.Budget

	if result := service.database.First(&budget, id); result.Error != nil {
		return entities.Budget{}, result.Error
	}

	return budget, nil
}

func (service *BudgetsServiceImpl) FindByName(name string) (entities.Budget, error) {
	var budget entities.Budget

	if result := service.database.First(&budget, "name = ?", name); result.Error != nil {
		return entities.Budget{}, result.Error
	}

	return budget, nil
}

func (service *BudgetsServiceImpl) Update(id string, updateBudgetDto UpdateBudgetDto) (entities.Budget, error) {
	var budget entities.Budget

	updatedBudget := entities.Budget{
		Name: updateBudgetDto.Name,
	}

	if result := service.database.First(&budget, id); result.Error != nil {
		return entities.Budget{}, result.Error
	}

	if result := service.database.Model(&budget).Updates(updatedBudget); result.Error != nil {
		return entities.Budget{}, result.Error
	}

	return budget, nil
}

func (service *BudgetsServiceImpl) Delete(id string) (entities.Budget, error) {
	var budget entities.Budget

	if result := service.database.First(&budget, id); result.Error != nil {
		return entities.Budget{}, result.Error
	}

	if result := service.database.Delete(&entities.Budget{}, id); result.Error != nil {
		return entities.Budget{}, result.Error
	}

	return budget, nil
}
