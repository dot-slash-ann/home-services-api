package services

import (
	"github.com/dot-slash-ann/home-services-api/database"
	"github.com/dot-slash-ann/home-services-api/dtos"
	"github.com/dot-slash-ann/home-services-api/entities"
)

func TransactionsFindAll() ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	results := database.Database.Find(&transactions)

	if results.Error != nil {
		return []entities.Transaction{}, results.Error

	}
	return transactions, nil
}

func TransactionsFindOne(id string) (entities.Transaction, error) {
	var transaction entities.Transaction

	results := database.Database.First(&transaction, id)

	if results.Error != nil {
		return entities.Transaction{}, results.Error
	}

	return transaction, nil
}

func TransactionsCreate(createTransactionDto dtos.CreateTransactionDto) (entities.Transaction, error) {
	transaction := entities.Transaction{
		TransactionOn: createTransactionDto.TransactionOn,
		PostedOn:      createTransactionDto.PostedOn,
		Amount:        createTransactionDto.Amount,
		VendorId:      createTransactionDto.VendorId,
		CategoryId:    createTransactionDto.CategoryId,
	}

	result := database.Database.Create(&transaction)

	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}

	return transaction, nil
}

func TransactionsUpdate(id string, updateTransactionDto dtos.UpdateTransactionDto) (entities.Transaction, error) {
	var transaction entities.Transaction

	result := database.Database.First(&transaction, id)

	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}

	result = database.Database.Model(&transaction).Updates(entities.Transaction{
		TransactionOn: updateTransactionDto.TransactionOn,
		PostedOn:      updateTransactionDto.PostedOn,
		Amount:        updateTransactionDto.Amount,
		CategoryId:    updateTransactionDto.CategoryId,
	})

	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}

	return transaction, nil
}

func TransactionsDelete(id string) (entities.Transaction, error) {
	var transaction entities.Transaction

	result := database.Database.First(&transaction, id)

	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}

	database.Database.Delete(&entities.Transaction{}, id)

	return transaction, nil
}
