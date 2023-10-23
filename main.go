package main

import (
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/controllers"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type RequestBody struct {
	Image string
}

func main() {
	godotenv.Load()

	db.Connect()

	router := gin.Default()

	router.GET("/healthcheck", controllers.HealthCheck())
	router.GET("/devices", middlewares.Authorize, controllers.GetAllDevices())
	router.POST("/devices", middlewares.Authorize, controllers.CreateDevice())
	router.GET("/devices/", middlewares.Authorize, controllers.GetDeviceByID())
	router.PUT("/devices/", middlewares.Authorize, controllers.UpdateDeviceByID())
	router.DELETE("/devices/", middlewares.Authorize, controllers.DeleteDeviceByID())
	router.GET("/firmwares", middlewares.Authorize, controllers.GetAllFirmwares())
	router.POST("/firmwares", middlewares.Authorize, controllers.CreateFirmware())
	router.GET("/firmwares/", middlewares.Authorize, controllers.GetFirmwareByID())
	router.PUT("/firmwares/", middlewares.Authorize, controllers.UpdateFirmwareByID())
	router.DELETE("/firmwares/", middlewares.Authorize, controllers.DeleteFirmwareByID())
	router.POST("/login", middlewares.Authorize, controllers.Login())
	router.POST("/signup", middlewares.Authorize, controllers.SignUp())

	router.Run()
}
