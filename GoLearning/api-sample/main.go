package main

import (
	"golearning/api-sample/controllers"
	"golearning/api-sample/repository"
	"golearning/api-sample/services"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "golearning/api-sample/docs"
)

var router = gin.Default()

// @title Boilerplate API in GO
// @version 1.0
// @description This is a sample RESTFULL API

// @contact.name GitHub
// @contact.url https://www.github.com/joemersonsouza

// @license.name BSD License

// @BasePath /v1
func main() {
	repo := repository.NewShopRepository()
	shopService := services.NewShopService(repo)
	controller := controllers.NewShopController(shopService)
	initRouter(controller)

	var db = repo.CreateRepositoryInstance()

	defer db.Close()

	router.Run("localhost:9090")

}

func initRouter(controller controllers.IShopController) {
	v1 := router.Group("v1")
	v1.GET("/items", controller.GetItems)
	v1.GET("/item/:id", controller.GetItem)
	v1.POST("/item", controller.AddItem)
	v1.PUT("/item/:id", controller.UpdateItem)
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
