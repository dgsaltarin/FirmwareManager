package controllers

import (
	"net/http"
	"time"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Create Device method
func CreateDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// bin user info from request
		var device models.Device
		if err := c.ShouldBindJSON(&device); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// add infor to device
		device.ID = uuid.New().String()
		device.CreatedAt = time.Now()
		device.UpdatedAt = time.Now()

		// create new device in database
		db.DB.Create(&device)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Devide created successfully",
		})
	}
}

// Get Device by ID method
func GetDeviceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameter value from the request
		deviceID := c.Query("id")
		var device models.Device

		// get users from database
		if err := db.DB.Where("id = ?", deviceID).First(&device).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "device not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"device": device,
		})
	}
}

// Get All Devices method
func GetAllDevices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var devices []models.Device

		// get users from database
		db.DB.Find(&devices)

		c.JSON(http.StatusOK, gin.H{
			"devices": devices,
		})
	}
}
