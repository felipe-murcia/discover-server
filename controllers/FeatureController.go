package controllers

import (
	"net/http"

	"discoverco.co/server/services"
	"github.com/gin-gonic/gin"
)

type FeatureController struct {
	service services.FeatureService
}

func NewFeatureController(service services.FeatureService) *FeatureController {
	return &FeatureController{service: service}
}

func (ctrl *FeatureController) FindAll(c *gin.Context) {
	categories, err := ctrl.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}
