package main

import (
	"flag"
	"fmt"
	"golearning/api-sample/controllers"
	"golearning/api-sample/repository"
	"golearning/api-sample/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "golearning/api-sample/docs"
	"golearning/api-sample/internal/config"
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
	var configFile string

	// -c options set configuration file path, but can be overwritten by GAPI_CONFIG_FILE environment variable
	flag.StringVar(&configFile, "c", "configs/dev.yaml", "config file path")
	flag.Parse()

	// If you specify an option by using environment variables, it overrides any value loaded from the configuration file
	path := os.Getenv("GAPI_CONFIG_FILE")
	if path != "" {
		configFile = path
	}

	// Load configuration yaml file using -c location/GAPI_CONFIG_FILE and merging environments variables with higher precedence
	sc, err := config.LoadServiceConfig(configFile)
	if err != nil {
		log.Fatalf("main: could not load service configuration [%v]", err)
	}

	repo := repository.NewShopRepository()
	shopService := services.NewShopService(repo)
	controller := controllers.NewShopController(shopService)

	repoConfig := repository.RepositoryConfig{
		Dialect:  sc.DB.Dialect,
		Host:     sc.DB.Host,
		DBport:   sc.DB.Port,
		User:     sc.DB.User,
		Database: sc.DB.Database,
		Password: sc.DB.Password,
	}

	fmt.Printf("Config: %v", repoConfig.Host)

	var db = repo.CreateRepositoryInstance(&repoConfig)

	defer db.Close()

	initRouter(controller)

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
