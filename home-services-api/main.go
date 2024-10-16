package main

import (
	"time"

	"github.com/dot-slash-ann/home-services-api/budgets"
	"github.com/dot-slash-ann/home-services-api/categories"
	"github.com/dot-slash-ann/home-services-api/database"
	"github.com/dot-slash-ann/home-services-api/initializers"
	"github.com/dot-slash-ann/home-services-api/middleware"
	"github.com/dot-slash-ann/home-services-api/tags"
	"github.com/dot-slash-ann/home-services-api/transactions"
	"github.com/dot-slash-ann/home-services-api/users"
	"github.com/dot-slash-ann/home-services-api/vendors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(middleware.ErrorHandler())

	vendorsService := vendors.NewVendorsService(database.Connection)
	vendorsController := vendors.NewVendorsController(vendorsService)

	router.POST("api/vendors", vendorsController.Create)
	router.GET("api/vendors", vendorsController.FindAll)
	router.GET("api/vendors/:id", vendorsController.FindOne)
	router.PATCH("api/vendors/:id", vendorsController.Update)
	router.DELETE("api/vendors/:id", vendorsController.Delete)

	categoriesService := categories.NewCategoriesService(database.Connection)
	categoriesController := categories.NewCategoriesController(categoriesService)

	router.POST("api/categories", categoriesController.Create)
	router.GET("api/categories", categoriesController.FindAll)
	router.GET("api/categories/:id", categoriesController.FindOne)
	router.PATCH("api/categories/:id", categoriesController.Update)
	router.DELETE("api/categories/:id", categoriesController.Delete)

	tagsService := tags.NewTagsService(database.Connection)
	tagsController := tags.NewTagsController(tagsService)

	router.POST("api/tags", tagsController.Create)
	router.GET("api/tags", tagsController.FindAll)
	router.GET("api/tags/:id", tagsController.FindOne)
	router.PATCH("api/tags/:id", tagsController.Update)
	router.DELETE("api/tags/:id", tagsController.Delete)

	transactionsService := transactions.NewTransactionsService(database.Connection, categoriesService, tagsService, vendorsService)
	transactionsController := transactions.NewTransactionsController(transactionsService)

	router.POST("api/transactions", transactionsController.Create)
	router.GET("api/transactions", transactionsController.FindAll)
	router.GET("api/transactions/:id", transactionsController.FindOne)
	router.PATCH("api/transactions/:id", transactionsController.Update)
	router.DELETE("api/transactions/:id", transactionsController.Delete)
	router.POST("api/transaction/:id/tag", transactionsController.TagTransaction)

	budgetsService := budgets.NewBudgetsService(database.Connection, categoriesService)
	budgetsController := budgets.NewBudgetsController(budgetsService)

	router.POST("api/budgets", budgetsController.Create)
	router.GET("api/budgets", budgetsController.FindAll)
	router.GET("api/budgets/:id", budgetsController.FindOne)
	router.DELETE("api/budgets/:id", budgetsController.Delete)
	router.POST("api/budgets/:id/category", budgetsController.AddCategory)

	usersService := users.NewUsersService(database.Connection)
	usersController := users.NewUsersController(usersService)

	router.POST("api/signup", usersController.SignUp)
	router.POST("api/login", usersController.Login)
	router.GET("api/users", middleware.RequireAuth(usersService), usersController.FindAll)
	router.GET("api/users/:id", middleware.RequireAuth(usersService), usersController.FindOne)

	router.Run()
}
