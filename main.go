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
	// router.PUT("/devices:id", controllers.CreateUser())
	// router.DELETE("/devices:id", controllers.CreateUser())
	// router.GET("/firmwares", controllers.UploadImage())
	// router.POST("/firmwares", controllers.CreateUser())
	// router.GET("/firmwares:id", controllers.UploadImage())
	// router.PUT("/firmwares:id", controllers.CreateUser())
	// router.DELETE("/firmwares:id", controllers.CreateUser())
	router.GET("/users", controllers.GetUserByID())
	router.POST("/users", controllers.CreateUser())

	router.Run()
}
