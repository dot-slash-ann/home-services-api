package main

import (
	CategoriesController "github.com/dot-slash-ann/home-services-api/controllers/categories"
	TagsController "github.com/dot-slash-ann/home-services-api/controllers/tags"
	TransactionsController "github.com/dot-slash-ann/home-services-api/controllers/transactions"
	"github.com/dot-slash-ann/home-services-api/database"
	"github.com/dot-slash-ann/home-services-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	router := gin.Default()

	router.POST("transactions", TransactionsController.Create)
	router.GET("transactions", TransactionsController.FindAll)
	router.GET("transactions/:id", TransactionsController.FindOne)
	router.PATCH("transactions/:id", TransactionsController.Update)
	router.DELETE("transactions/:id", TransactionsController.Delete)

	router.POST("categories", CategoriesController.Create)
	router.GET("categories", CategoriesController.FindAll)
	router.GET("categories/:id", CategoriesController.FindOne)
	router.PATCH("categories/:id", CategoriesController.Update)
	router.DELETE("categories/:id", CategoriesController.Delete)

	router.POST("tags", TagsController.Create)
	router.GET("tags", TagsController.FindAll)
	router.GET("tags/:id", TagsController.FindOne)
	router.PATCH("tags/:id", TagsController.Update)
	router.DELETE("tags/:id", TagsController.Delete)

	router.Run()
}
