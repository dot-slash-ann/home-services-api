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

	router.GET("transactions", controllers.FindAll)
	router.GET("transactions/:id", controllers.FindOne)
	router.POST("transactions", controllers.Create)
	router.PUT("transactions/:id", controllers.Update)
	router.DELETE("transactions/:id", controllers.Delete)

	router.Run()

	// result := countIslands(matrix)
}
