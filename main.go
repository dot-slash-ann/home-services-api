package main

import (
	"github.com/dot-slash-ann/home-services-api/controllers"
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

	router.GET("transactions", controllers.FindAllTransactions)
	router.GET("transactions/:id", controllers.FindOneTransaction)
	router.POST("transactions", controllers.CreateTransaction)
	router.PUT("transactions/:id", controllers.UpdateTransaction)
	router.DELETE("transactions/:id", controllers.DeleteTransaction)

	router.Run()
}
