package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"restapi/models"
	"restapi/services"
)

type LocationController struct {
	LocationService services.LocationService

}

func NewLocationController(locationService services.LocationService) LocationController {
	return LocationController {
		LocationService: locationService,
	}
}

func (lc *LocationController) GetAll(ctx *gin.Context) {
	locations, err := lc.LocationService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, locations)
}

func (lc *LocationController) Add(ctx *gin.Context) {
	var location models.Location
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := lc.LocationService.Add(&location)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (lc *LocationController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var location models.Location
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := lc.LocationService.Update(&location, &id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (lc *LocationController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := lc.LocationService.Delete(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaage": "success"})
}

func (lc *LocationController) RegisterRoutes(rg *gin.RouterGroup) {
	locationGroup := rg.Group("/location")
	locationGroup.GET("/", lc.GetAll)
	locationGroup.POST("/", lc.Add)
	locationGroup.PUT("/:id", lc.Update)
	locationGroup.DELETE("/:id", lc.Delete)
}