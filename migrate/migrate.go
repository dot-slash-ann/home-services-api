package main

import (
	"github.com/dot-slash-ann/home-services-api/database"
	CategoriesEntity "github.com/dot-slash-ann/home-services-api/entities/categories"
	TransactionsEntity "github.com/dot-slash-ann/home-services-api/entities/transactions"
	"github.com/dot-slash-ann/home-services-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	database.Connection.AutoMigrate(&TransactionsEntity.Transaction{})
	database.Connection.AutoMigrate(&CategoriesEntity.Category{})
}
