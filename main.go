package main

import (
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/controllers"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
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
	router.GET("/devices", controllers.GetAllDevices())
	router.POST("/devices", controllers.CreateDevice())
	router.GET("/devices:id", controllers.GetDeviceByID())
	router.PUT("/devices:id", controllers.UpdateDeviceByID())
	router.DELETE("/devices:id", controllers.DeleteDeviceByID())
	router.GET("/firmwares", controllers.GetAllFirmwares())
	router.POST("/firmwares", controllers.CreateFirmware())
	router.GET("/firmwares:id", controllers.GetFirmwareByID())
	router.PUT("/firmwares:id", controllers.UpdateFirmwareByID())
	router.DELETE("/firmwares:id", controllers.DeleteFirmwareByID())
	router.POST("/login", controllers.Login())
	router.POST("/signup", controllers.SignUp())

	router.Run()
}
