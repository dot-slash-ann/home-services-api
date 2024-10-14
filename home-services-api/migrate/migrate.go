package main

import (
	"github.com/dot-slash-ann/home-services-api/database"
	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/dot-slash-ann/home-services-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	database.Connection.AutoMigrate(&entities.Vendor{})
	database.Connection.AutoMigrate(&entities.Category{})
	database.Connection.AutoMigrate(&entities.Tag{})
	database.Connection.AutoMigrate(&entities.Transaction{})

	database.Connection.AutoMigrate(&entities.Budget{})
	database.Connection.AutoMigrate(&entities.BudgetCategory{})

	database.Connection.AutoMigrate(&entities.User{})
}
