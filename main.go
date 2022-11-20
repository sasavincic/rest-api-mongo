package main

import (
	"context"
	"log"

	"restapi/utils"
	"restapi/services"
	"restapi/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server *gin.Engine

	carService services.CarService
	carController controllers.CarController
	carCollection *mongo.Collection

	chargingStationService services.ChargingStationService
	chargingStationController controllers.ChargingStationController
	chargingStationCollection *mongo.Collection

	locationService services.LocationService
	locationController controllers.LocationController
	locationCollection *mongo.Collection

	ctx context.Context
	mongoclient *mongo.Client
	err error
)

func init() {
	ctx = context.TODO()
	mongoclient := utils.InitDB(ctx)

	carCollection = mongoclient.Database("ev_route_planner").Collection("cars")
	carService = services.NewCarService(carCollection, ctx)
	carController = controllers.NewCarController(carService)

	chargingStationCollection = mongoclient.Database("ev_route_planner").Collection("chargingstations")
	chargingStationService = services.NewChargingStationService(chargingStationCollection, ctx)
	chargingStationController = controllers.NewChargingStationController(chargingStationService)

	locationCollection = mongoclient.Database("ev_route_planner").Collection("locations")
	locationService = services.NewLocationService(locationCollection, ctx)
	locationController = controllers.NewLocationController(locationService)
	
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("api/v1")

	carController.RegisterRoutes(basepath)
	chargingStationController.RegisterRoutes(basepath)
	locationController.RegisterRoutes(basepath)

	log.Fatal(server.Run(":8080"))
}