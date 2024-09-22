package TransactionsService

import (
	"github.com/dot-slash-ann/home-services-api/database"
	TransactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	TransactionsEntity "github.com/dot-slash-ann/home-services-api/entities/transactions"
)

func TransactionsFindAll() ([]TransactionsEntity.Transaction, error) {
	var transactions []TransactionsEntity.Transaction

	results := database.Database.Find(&transactions)

	if results.Error != nil {
		return []TransactionsEntity.Transaction{}, results.Error

	}
	return transactions, nil
}

func TransactionsFindOne(id string) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	results := database.Database.First(&transaction, id)

	if results.Error != nil {
		return TransactionsEntity.Transaction{}, results.Error
	}

	return transaction, nil
}

func TransactionsCreate(createTransactionDto TransactionsDto.CreateTransactionDto) (TransactionsEntity.Transaction, error) {
	transaction := TransactionsEntity.Transaction{
		TransactionOn: createTransactionDto.TransactionOn,
		PostedOn:      createTransactionDto.PostedOn,
		Amount:        createTransactionDto.Amount,
		VendorId:      createTransactionDto.VendorId,
		CategoryId:    createTransactionDto.CategoryId,
	}

	result := database.Database.Create(&transaction)

	if result.Error != nil {
		return TransactionsEntity.Transaction{}, result.Error
	}

	return transaction, nil
}

func TransactionsUpdate(id string, updateTransactionDto TransactionsDto.UpdateTransactionDto) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	result := database.Database.First(&transaction, id)

	if result.Error != nil {
		return TransactionsEntity.Transaction{}, result.Error
	}

	result = database.Database.Model(&transaction).Updates(TransactionsEntity.Transaction{
		TransactionOn: updateTransactionDto.TransactionOn,
		PostedOn:      updateTransactionDto.PostedOn,
		Amount:        updateTransactionDto.Amount,
		CategoryId:    updateTransactionDto.CategoryId,
	})

	if result.Error != nil {
		return TransactionsEntity.Transaction{}, result.Error
	}

	return transaction, nil
}

func TransactionsDelete(id string) (TransactionsEntity.Transaction, error) {
	var transaction TransactionsEntity.Transaction

	result := database.Database.First(&transaction, id)

	if result.Error != nil {
		return TransactionsEntity.Transaction{}, result.Error
	}

	database.Database.Delete(&TransactionsEntity.Transaction{}, id)

	return transaction, nil
}
