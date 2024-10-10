package main

import (
	categoriesController "github.com/dot-slash-ann/home-services-api/controllers/categories"
	tagsController "github.com/dot-slash-ann/home-services-api/controllers/tags"
	transactionsController "github.com/dot-slash-ann/home-services-api/controllers/transactions"
	usersController "github.com/dot-slash-ann/home-services-api/controllers/users"
	vendorsController "github.com/dot-slash-ann/home-services-api/controllers/vendors"
	"github.com/dot-slash-ann/home-services-api/database"
	"github.com/dot-slash-ann/home-services-api/initializers"
	"github.com/dot-slash-ann/home-services-api/middleware"
	"github.com/dot-slash-ann/home-services-api/services/categories"
	"github.com/dot-slash-ann/home-services-api/services/tags"
	"github.com/dot-slash-ann/home-services-api/services/transactions"
	"github.com/dot-slash-ann/home-services-api/services/users"
	"github.com/dot-slash-ann/home-services-api/services/vendors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	router := gin.Default()

	router.Use(middleware.ErrorHandler())

	vendorsService := vendors.NewVendorsService(database.Connection)
	vendorsController := vendorsController.NewVendorsController(vendorsService)

	router.POST("api/vendors", vendorsController.Create)
	router.GET("api/vendors", vendorsController.FindAll)
	router.GET("api/vendors/:id", vendorsController.FindOne)
	router.PATCH("api/vendors/:id", vendorsController.Update)
	router.DELETE("api/vendors/:id", vendorsController.Delete)

	categoriesService := categories.NewCategoriesService(database.Connection)
	categoriesController := categoriesController.NewCategoriesController(categoriesService)

	router.POST("api/categories", categoriesController.Create)
	router.GET("api/categories", categoriesController.FindAll)
	router.GET("api/categories/:id", categoriesController.FindOne)
	router.PATCH("api/categories/:id", categoriesController.Update)
	router.DELETE("api/categories/:id", categoriesController.Delete)

	tagsService := tags.NewTagsService(database.Connection)
	tagsController := tagsController.NewTagsController(tagsService)

	router.POST("api/tags", tagsController.Create)
	router.GET("api/tags", tagsController.FindAll)
	router.GET("api/tags/:id", tagsController.FindOne)
	router.PATCH("api/tags/:id", tagsController.Update)
	router.DELETE("api/tags/:id", tagsController.Delete)

	transactionsService := transactions.NewTransactionsService(database.Connection, categoriesService, tagsService)
	transactionsController := transactionsController.NewTransactionsController(transactionsService)

	router.POST("api/transactions", transactionsController.Create)
	router.GET("api/transactions", transactionsController.FindAll)
	router.GET("api/transactions/:id", transactionsController.FindOne)
	router.PATCH("api/transactions/:id", transactionsController.Update)
	router.DELETE("api/transactions/:id", transactionsController.Delete)
	router.POST("api/transaction/:id/tag", transactionsController.TagTransaction)

	usersService := users.NewUsersService(database.Connection)
	usersController := usersController.NewUsersController(usersService)

	router.POST("api/signup", usersController.SignUp)
	router.POST("api/login", usersController.Login)
	router.GET("api/users", middleware.RequireAuth(usersService), usersController.FindAll)
	router.GET("api/users/:id", middleware.RequireAuth(usersService), usersController.FindOne)

	router.Run()
}
