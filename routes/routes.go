package routes

import (
	"discoverco.co/server/controllers"
	"discoverco.co/server/repositories"
	"discoverco.co/server/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	featureRepository := repositories.NewFeatureRepositoryImpl()
	featureService := services.NewFeatureService(featureRepository)
	featureController := controllers.NewFeatureController(featureService)

	router.GET("/features", featureController.FindAll)

	return router
}
