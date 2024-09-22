package main

import (
	"github.com/dot-slash-ann/home-services-api/database"
	TransactionsEntity "github.com/dot-slash-ann/home-services-api/entities/transactions"
	"github.com/dot-slash-ann/home-services-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	database.Database.AutoMigrate(&TransactionsEntity.Transaction{})
}
