package main

import (
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/controllers"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Image string
}

func main() {
	router := gin.Default()

	router.GET("/healthcheck", controllers.HealthCheck())
	router.GET("/devices", controllers.CreateDevice())
	router.POST("/devices", controllers.CreateDevice())
	router.GET("/devices", controllers.GetDeviceByID())
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
