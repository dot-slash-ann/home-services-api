package transactions

import (
	"fmt"
	"strings"

	transactionsDto "github.com/dot-slash-ann/home-services-api/dtos/transactions"
	"github.com/dot-slash-ann/home-services-api/entities/transactions"
	"github.com/dot-slash-ann/home-services-api/services/categories"
	"github.com/dot-slash-ann/home-services-api/services/tags"
	"github.com/dot-slash-ann/home-services-api/services/vendors"
	"gorm.io/gorm"
)

type TransactionsService interface {
	Create(transactionsDto.CreateTransactionDto) (transactions.Transaction, error)
	FindAll(map[string]string) ([]transactions.Transaction, error)
	FindOne(string) (transactions.Transaction, error)
	Update(string, transactionsDto.UpdateTransactionDto) (transactions.Transaction, error)
	Delete(string) (transactions.Transaction, error)
	TagTransaction(transactionsDto.TagTransactionDto, string) (transactions.Transaction, error)
}

type TransactionsServiceImpl struct {
	database          *gorm.DB
	categoriesService categories.CategoriesService
	tagsService       tags.TagsService
	vendorsService    vendors.VendorsService
}

func NewTransactionsService(database *gorm.DB, categoriesService categories.CategoriesService, tagsService tags.TagsService, vendorsService vendors.VendorsService) *TransactionsServiceImpl {
	return &TransactionsServiceImpl{
		database:          database,
		categoriesService: categoriesService,
		tagsService:       tagsService,
		vendorsService:    vendorsService,
	}
}

func (service *TransactionsServiceImpl) Create(createTransactionDto transactionsDto.CreateTransactionDto) (transactions.Transaction, error) {
	category, err := service.categoriesService.FindByName(createTransactionDto.CategoryName)

	if err != nil {
		return transactions.Transaction{}, err
	}

	vendor, err := service.vendorsService.FindByName(createTransactionDto.VendorName)

	if err != nil {
		return transactions.Transaction{}, err
	}

	transaction := transactions.Transaction{
		TransactionOn:   createTransactionDto.TransactionOn,
		PostedOn:        createTransactionDto.PostedOn,
		Amount:          createTransactionDto.Amount,
		TransactionType: createTransactionDto.TransactionType,
		VendorID:        vendor.ID,
		CategoryID:      category.ID,
	}

	if result := service.database.Create(&transaction); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	if result := service.database.Preload("Category").Preload("Tags").Preload("Vendor").First(&transaction, fmt.Sprint(transaction.ID)); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) FindAll(filters map[string]string) ([]transactions.Transaction, error) {
	var transactionsList []transactions.Transaction

	query := service.database.Model(&transactions.Transaction{}).Preload("Category").Preload("Tags").Preload("Vendor")

	if categoryName, ok := filters["categoryName"]; ok {
		category, err := service.categoriesService.FindByName(categoryName)

		if err != nil {
			return []transactions.Transaction{}, err
		}

		query.Where("category_id = ?", category.ID)
	} else if categoryID, ok := filters["categoryID"]; ok {
		category, err := service.categoriesService.FindOne(categoryID)

		if err != nil {
			return []transactions.Transaction{}, err
		}

		query.Where("category_id = ?", category.ID)
	}

	if vendorID, ok := filters["vendorID"]; ok {
		vendor, err := service.vendorsService.FindOne(vendorID)

		if err != nil {
			return []transactions.Transaction{}, err
		}

		query.Where("vendor_id = ?", vendor.ID)
	}

	if tags, ok := filters["tags"]; ok && tags != "" {
		tagList := strings.Split(tags, ",")

		query = query.Joins("JOIN transaction_tags ON transactions.id = transaction_tags.transaction_id").
			Joins("JOIN tags on transaction_tags.tag_id = tags.id").
			Where("tags.name IN (?)", tagList)
	}

	if min, ok := filters["min"]; ok && min != "" {
		query.Where("transactions.amount >= ?", min)
	}

	if max, ok := filters["max"]; ok && max != "" {
		query.Where("transactions.amount <= ?", max)
	}

	if transactionOnFrom, ok := filters["transactionOnFrom"]; ok && transactionOnFrom != "" {
		query.Where("transactions.transaction_on >= ?", transactionOnFrom)
	}

	if transactionOnTo, ok := filters["transactionOnTo"]; ok && transactionOnTo != "" {
		query.Where("transactions.transaction_on <= ?", transactionOnTo)
	}

	if postedOnFrom, ok := filters["postedOnFrom"]; ok && postedOnFrom != "" {
		query.Where("transactions.posted_on >= ?", postedOnFrom)
	}

	if postedOnTo, ok := filters["postedOnTo"]; ok && postedOnTo != "" {
		query.Where("transactions.posted_on <= ?", postedOnTo)
	}

	query.Order("transaction_on ASC").Order("posted_on ASC")
	query.Limit(100)

	if result := query.Find(&transactionsList); result.Error != nil {
		return []transactions.Transaction{}, result.Error
	}

	if len(transactionsList) == 0 {
		return []transactions.Transaction{}, nil
	}

	return transactionsList, nil
}

func (service *TransactionsServiceImpl) FindOne(id string) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if result := service.database.Preload("Category").Preload("Vendor").Preload("Tags").First(&transaction, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) Update(id string, updateTransactionDto transactionsDto.UpdateTransactionDto) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if result := service.database.First(&transaction, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	category, err := service.categoriesService.FindOne(fmt.Sprint(updateTransactionDto.CategoryId))

	if err != nil && updateTransactionDto.CategoryId != 0 {
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

	if result := service.database.Preload("Category").Preload("Tags").Preload("Vendor").First(&transaction, fmt.Sprint(transaction.ID)); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) Delete(id string) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if result := service.database.Preload("Category").Preload("Tags").Preload("Vendor").First(&transaction, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	if result := service.database.Delete(&transactions.Transaction{}, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return transaction, nil
}

func (service *TransactionsServiceImpl) TagTransaction(tagTransactionDto transactionsDto.TagTransactionDto, id string) (transactions.Transaction, error) {
	var transaction transactions.Transaction

	if result := service.database.Preload("Category").Preload("Tags").Preload("Vendor").First(&transaction, id); result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	tag, err := service.tagsService.FindOneOrCreate(tagTransactionDto.TagName)

	if err != nil {
		return transactions.Transaction{}, err
	}

	if err := service.database.Model(&transaction).Association("Tags").Append(&tag); err != nil {
		return transactions.Transaction{}, err
	}

	return transaction, nil
}
