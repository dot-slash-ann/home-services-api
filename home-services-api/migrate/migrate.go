package main

import (
	"github.com/dot-slash-ann/home-services-api/database"
	"github.com/dot-slash-ann/home-services-api/entities/categories"
	"github.com/dot-slash-ann/home-services-api/entities/tags"
	"github.com/dot-slash-ann/home-services-api/entities/transactions"
	"github.com/dot-slash-ann/home-services-api/entities/users"
	"github.com/dot-slash-ann/home-services-api/entities/vendors"
	"github.com/dot-slash-ann/home-services-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	database.Connection.AutoMigrate(&vendors.Vendor{})
	database.Connection.AutoMigrate(&categories.Category{})
	database.Connection.AutoMigrate(&tags.Tag{})
	database.Connection.AutoMigrate(&transactions.Transaction{})

	database.Connection.AutoMigrate(&users.User{})
}
