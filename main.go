package main

import (
	CategoriesController "github.com/dot-slash-ann/home-services-api/controllers/categories"
	TagsController "github.com/dot-slash-ann/home-services-api/controllers/tags"
	TransactionsController "github.com/dot-slash-ann/home-services-api/controllers/transactions"
	UsersController "github.com/dot-slash-ann/home-services-api/controllers/users"
	"github.com/dot-slash-ann/home-services-api/database"
	"github.com/dot-slash-ann/home-services-api/initializers"
	"github.com/dot-slash-ann/home-services-api/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	router := gin.Default()

	router.POST("api/transactions", TransactionsController.Create)
	router.GET("api/transactions", TransactionsController.FindAll)
	router.GET("api/transactions/:id", TransactionsController.FindOne)
	router.PATCH("api/transactions/:id", TransactionsController.Update)
	router.DELETE("api/transactions/:id", TransactionsController.Delete)

	router.POST("api/categories", CategoriesController.Create)
	router.GET("api/categories", CategoriesController.FindAll)
	router.GET("api/categories/:id", CategoriesController.FindOne)
	router.PATCH("api/categories/:id", CategoriesController.Update)
	router.DELETE("api/categories/:id", CategoriesController.Delete)

	router.POST("api/tags", TagsController.Create)
	router.GET("api/tags", TagsController.FindAll)
	router.GET("api/tags/:id", TagsController.FindOne)
	router.PATCH("api/tags/:id", TagsController.Update)
	router.DELETE("api/tags/:id", TagsController.Delete)

	router.POST("api/signup", UsersController.SignUp)
	router.POST("api/login", UsersController.Login)
	router.GET("api/users", middleware.RequireAuth, UsersController.FindAll)
	router.GET("api/users/:id", middleware.RequireAuth, UsersController.FindOne)

	router.Run()
}
