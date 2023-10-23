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
	router.GET("/devices/", controllers.GetDeviceByID())
	router.PUT("/devices/", controllers.UpdateDeviceByID())
	router.DELETE("/devices/", controllers.DeleteDeviceByID())
	router.GET("/firmwares", controllers.GetAllFirmwares())
	router.POST("/firmwares", controllers.CreateFirmware())
	router.GET("/firmwares/", controllers.GetFirmwareByID())
	router.PUT("/firmwares/", controllers.UpdateFirmwareByID())
	router.DELETE("/firmwares/", controllers.DeleteFirmwareByID())
	router.POST("/login", controllers.Login())
	router.POST("/signup", controllers.SignUp())

	router.Run()
}
