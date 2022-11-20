package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"restapi/models"
	"restapi/services"
)

type CarController struct {
	CarService services.CarService

}

func NewCarController(carService services.CarService) CarController {
	return CarController {
		CarService: carService,
	}
}

func (cc *CarController) GetAll(ctx *gin.Context) {
	cars, err := cc.CarService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cars)
}

func (cc *CarController) Add(ctx *gin.Context) {
	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := cc.CarService.Add(&car)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (cc *CarController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := cc.CarService.Update(&car, &id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (cc *CarController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := cc.CarService.Delete(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (cc *CarController) RegisterRoutes(rg *gin.RouterGroup) {
	chargingStationGroup := rg.Group("/car")
	chargingStationGroup.GET("/", cc.GetAll)
	chargingStationGroup.POST("/", cc.Add)
	chargingStationGroup.PUT("/:id", cc.Update)
	chargingStationGroup.DELETE("/:id", cc.Delete)
}