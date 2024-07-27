package routes

import (
	"discoverco.co/server/controllers"
	"discoverco.co/server/repositories"
	"discoverco.co/server/services"
	"github.com/gin-gonic/gin"
)

func FeatureRouter(router *gin.Engine) {
	routes := router.Group("api/v1/features")

	featureRepository := repositories.NewFeatureRepositoryImpl()
	featureService := services.NewFeatureService(featureRepository)
	featureController := controllers.NewFeatureController(featureService)

	routes.GET("", featureController.FindAll)

}
