package transactions

import (
	"fmt"

	TransactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	"github.com/dot-slash-ann/home-services-api/lib"
	"gorm.io/gorm"

	"github.com/dot-slash-ann/home-services-api/entities/transactions"

	"github.com/dot-slash-ann/home-services-api/services/categories"
)

type TransactionsService interface {
	Create(TransactionsDto.CreateTransactionDto) (transactions.Transaction, error)
	FindAll() ([]transactions.Transaction, error)
	FindOne(string) (transactions.Transaction, error)
	Update(string, TransactionsDto.UpdateTransactionDto) (transactions.Transaction, error)
	Delete(string) (transactions.Transaction, error)
}

type TransactionsServiceImpl struct {
	database          *gorm.DB
	categoriesService categories.CategoriesService
}

func NewTransactionsService(database *gorm.DB, categoriesService categories.CategoriesService) *TransactionsServiceImpl {
	return &TransactionsServiceImpl{
		database:          database,
		categoriesService: categoriesService,
	}
}

// func preloadTransaction(service *TransactionsServiceImpl, transaction *transactions.Transaction, id string) error {
// 	return lib.HandleDatabaseError(service.database.Preload("Category").First(transaction, id))
// }

// func findCategory(categoryId uint) (CategoriesEntity.Category, error) {
// 	categoriesService := categories.NewCategoriesService(*&gorm.DB{})

// 	category, err := categoriesService.FindOne(fmt.Sprint(categoryId))

// 	if err != nil {
// 		return CategoriesEntity.Category{}, err
// 	}

// 	return category, nil
// }

func (service *TransactionsServiceImpl) Create(createTransactionDto TransactionsDto.CreateTransactionDto) (transactions.Transaction, error) {
	category, err := service.FindOne(fmt.Sprint(createTransactionDto.CategoryId))

	if err != nil {
		return transactions.Transaction{}, err
	}

	transaction := transactions.Transaction{
		TransactionOn: createTransactionDto.TransactionOn,
		PostedOn:      createTransactionDto.PostedOn,
		Amount:        createTransactionDto.Amount,
		VendorId:      createTransactionDto.VendorId,
		CategoryID:    category.ID,
	}

	if err := lib.HandleDatabaseError(service.database.Create(&transaction)); err != nil {
		return transactions.Transaction{}, err
	}

	if err := lib.HandleDatabaseError(service.database.Preload("Category").First(&transaction, fmt.Sprint(transaction.ID))); err != nil {
		return transactions.Transaction{}, err
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) FindAll() ([]transactions.Transaction, error) {
	var transactionsList []transactions.Transaction

	if err := lib.HandleDatabaseError(service.database.Model(&transactions.Transaction{}).Preload("Category").Find(&transactionsList)); err != nil {
		return []transactions.Transaction{}, err

	}
	return transactionsList, nil
}

func (service *TransactionsServiceImpl) FindOne(id string) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if err := lib.HandleDatabaseError(service.database.Preload("Category").First(&transaction, id)); err != nil {
		return transactions.Transaction{}, err
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) Update(id string, updateTransactionDto TransactionsDto.UpdateTransactionDto) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if err := lib.HandleDatabaseError(service.database.First(&transaction, id)); err != nil {
		return transactions.Transaction{}, err
	}

	categoriesService := categories.NewCategoriesService(service.database)

	category, err := categoriesService.FindOne(fmt.Sprint(updateTransactionDto.CategoryId))

	if err != nil {
		return transactions.Transaction{}, err
	}

	updatedTransaction := transactions.Transaction{
		TransactionOn: updateTransactionDto.TransactionOn,
		PostedOn:      updateTransactionDto.PostedOn,
		Amount:        updateTransactionDto.Amount,
		CategoryID:    category.ID,
	}

	if err := lib.HandleDatabaseError(service.database.Model(&transaction).Updates(updatedTransaction)); err != nil {
		return transactions.Transaction{}, err
	}

	if err := lib.HandleDatabaseError(service.database.Preload("Category").First(&transaction, fmt.Sprint(transaction.ID))); err != nil {
		return transactions.Transaction{}, err
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) Delete(id string) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if err := lib.HandleDatabaseError(service.database.Preload("Category").First(&transaction, id)); err != nil {
		return transactions.Transaction{}, err
	}

	if err := lib.HandleDatabaseError(service.database.Delete(&transactions.Transaction{}, id)); err != nil {
		return transactions.Transaction{}, err
	}

	return transaction, nil
}
