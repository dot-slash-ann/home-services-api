package TransactionsService

import (
	"fmt"

	"github.com/dot-slash-ann/home-services-api/database"
	TransactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	"gorm.io/gorm"

	CategoriesEntity "github.com/dot-slash-ann/home-services-api/entities/categories"
	TransactionsEntity "github.com/dot-slash-ann/home-services-api/entities/transactions"

	CategoriesService "github.com/dot-slash-ann/home-services-api/services/categories"
)

func handleDatabaseError(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func preloadTransaction(transaction *TransactionsEntity.Transaction, id string) error {
	return handleDatabaseError(database.Connection.Preload("Category").First(transaction, id))
}

func findCategory(categoryId uint) (CategoriesEntity.Category, error) {
	category, err := CategoriesService.FindOne(fmt.Sprint(categoryId))

	if err != nil {
		return CategoriesEntity.Category{}, err
	}

	return category, nil
}

func Create(createTransactionDto TransactionsDto.CreateTransactionDto) (TransactionsEntity.Transaction, error) {
	category, err := CategoriesService.FindOne(fmt.Sprint(createTransactionDto.CategoryId))

	if err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	transaction := TransactionsEntity.Transaction{
		TransactionOn: createTransactionDto.TransactionOn,
		PostedOn:      createTransactionDto.PostedOn,
		Amount:        createTransactionDto.Amount,
		VendorId:      createTransactionDto.VendorId,
		CategoryID:    category.ID,
	}

	if err := handleDatabaseError(database.Connection.Create(&transaction)); err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	if err := preloadTransaction(&transaction, fmt.Sprint(transaction.ID)); err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	return transaction, nil
}

func FindAll() ([]TransactionsEntity.Transaction, error) {
	var transactions []TransactionsEntity.Transaction

	if err := handleDatabaseError(database.Connection.Model(&TransactionsEntity.Transaction{}).Preload("Category").Find(&transactions)); err != nil {
		return []TransactionsEntity.Transaction{}, err

	}
	return transactions, nil
}

func FindOne(id string) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	if err := preloadTransaction(&transaction, id); err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	return transaction, nil
}

func Update(id string, updateTransactionDto TransactionsDto.UpdateTransactionDto) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	if err := handleDatabaseError(database.Connection.First(&transaction, id)); err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	category, err := findCategory(updateTransactionDto.CategoryId)
	if err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	updatedTransaction := TransactionsEntity.Transaction{
		TransactionOn: updateTransactionDto.TransactionOn,
		PostedOn:      updateTransactionDto.PostedOn,
		Amount:        updateTransactionDto.Amount,
		CategoryID:    category.ID,
	}

	if err := handleDatabaseError(database.Connection.Model(&transaction).Updates(updatedTransaction)); err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	if err := preloadTransaction(&transaction, fmt.Sprint(transaction.ID)); err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	return transaction, nil
}

func Delete(id string) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	if err := preloadTransaction(&transaction, id); err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	if err := handleDatabaseError(database.Connection.Delete(&TransactionsEntity.Transaction{}, id)); err != nil {
		return TransactionsEntity.Transaction{}, err
	}

	return transaction, nil
}
