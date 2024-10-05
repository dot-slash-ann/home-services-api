package transactions

import (
	"fmt"

	TransactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	"github.com/dot-slash-ann/home-services-api/entities/transactions"
	"github.com/dot-slash-ann/home-services-api/services/categories"
	"gorm.io/gorm"
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

	if result := service.database.Create(&transaction); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	if result := service.database.Preload("Category").First(&transaction, fmt.Sprint(transaction.ID)); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) FindAll() ([]transactions.Transaction, error) {
	var transactionsList []transactions.Transaction

	if result := service.database.Model(&transactions.Transaction{}).Preload("Category").Find(&transactionsList); result.Error != nil {
		return []transactions.Transaction{}, result.Error

	}
	return transactionsList, nil
}

func (service *TransactionsServiceImpl) FindOne(id string) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if result := service.database.Preload("Category").First(&transaction, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) Update(id string, updateTransactionDto TransactionsDto.UpdateTransactionDto) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if result := service.database.First(&transaction, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
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

	if result := service.database.Model(&transaction).Updates(updatedTransaction); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	if result := service.database.Preload("Category").First(&transaction, fmt.Sprint(transaction.ID)); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) Delete(id string) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if result := service.database.Preload("Category").First(&transaction, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	if result := service.database.Delete(&transactions.Transaction{}, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return transaction, nil
}
