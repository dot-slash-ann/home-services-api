package TransactionsService

import (
	"github.com/dot-slash-ann/home-services-api/database"
	TransactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	TransactionsEntity "github.com/dot-slash-ann/home-services-api/entities/transactions"
)

func Create(createTransactionDto TransactionsDto.CreateTransactionDto) (TransactionsEntity.Transaction, error) {
	transaction := TransactionsEntity.Transaction{
		TransactionOn: createTransactionDto.TransactionOn,
		PostedOn:      createTransactionDto.PostedOn,
		Amount:        createTransactionDto.Amount,
		VendorId:      createTransactionDto.VendorId,
		CategoryId:    createTransactionDto.CategoryId,
	}

	result := database.Connection.Create(&transaction)

	if result.Error != nil {
		return TransactionsEntity.Transaction{}, result.Error
	}

	return transaction, nil
}

func FindAll() ([]TransactionsEntity.Transaction, error) {
	var transactions []TransactionsEntity.Transaction

	results := database.Connection.Find(&transactions)

	if results.Error != nil {
		return []TransactionsEntity.Transaction{}, results.Error

	}
	return transactions, nil
}

func FindOne(id string) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	if results := database.Connection.First(&transaction, id); results.Error != nil {
		return TransactionsEntity.Transaction{}, results.Error
	}

	return transaction, nil
}

func Update(id string, updateTransactionDto TransactionsDto.UpdateTransactionDto) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	if result := database.Connection.First(&transaction, id); result.Error != nil {
		return TransactionsEntity.Transaction{}, result.Error
	}

	if result := database.Connection.Model(&transaction).Updates(TransactionsEntity.Transaction{
		TransactionOn: updateTransactionDto.TransactionOn,
		PostedOn:      updateTransactionDto.PostedOn,
		Amount:        updateTransactionDto.Amount,
		CategoryId:    updateTransactionDto.CategoryId,
	}); result.Error != nil {
		return TransactionsEntity.Transaction{}, result.Error
	}

	return transaction, nil
}

func Delete(id string) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	if result := database.Connection.First(&transaction, id); result.Error != nil {
		return TransactionsEntity.Transaction{}, result.Error
	}

	database.Connection.Delete(&TransactionsEntity.Transaction{}, id)

	return transaction, nil
}
