package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"restapi/models"
	"restapi/services"
)

type ChargingStationController struct {
	ChargingStationService services.ChargingStationService

}

func NewChargingStationController(chargingStationService services.ChargingStationService) ChargingStationController {
	return ChargingStationController {
		ChargingStationService: chargingStationService,
	}
}

func (csc *ChargingStationController) GetAll(ctx *gin.Context) {
	chargingstations, err := csc.ChargingStationService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, chargingstations)
}

func (csc *ChargingStationController) Add(ctx *gin.Context) {
	var chargingStation models.ChargingStation
	if err := ctx.ShouldBindJSON(&chargingStation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := csc.ChargingStationService.Add(&chargingStation)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (csc *ChargingStationController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var chargingStation models.ChargingStation
	if err := ctx.ShouldBindJSON(&chargingStation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := csc.ChargingStationService.Update(&chargingStation, &id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (csc *ChargingStationController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := csc.ChargingStationService.Delete(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (csc *ChargingStationController) RegisterRoutes(rg *gin.RouterGroup) {
	chargingStationGroup := rg.Group("/chargingstation")
	chargingStationGroup.GET("/", csc.GetAll)
	chargingStationGroup.POST("/", csc.Add)
	chargingStationGroup.PUT("/:id", csc.Update)
	chargingStationGroup.DELETE("/:id", csc.Delete)
}