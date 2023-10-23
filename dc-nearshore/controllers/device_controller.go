package controllers

import (
	"net/http"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/gin-gonic/gin"
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

		// create new device in database
		device, err := db.CreateDevice(device)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error creating device",
			})
			return
		}

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
		if _, err := db.GetDeviceByID(deviceID); err != nil {
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
		if _, err := db.GetAllDevices(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "devices not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"devices": devices,
		})
	}
}

// Delete Device by ID method
func DeleteDeviceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameter value from the request
		deviceID := c.Query("id")

		//check if device exists
		if _, err := db.GetDeviceByID(deviceID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "device not found",
			})
			return
		}

		// delete device from database
		db.DeleteDeviceByID(deviceID)

		c.JSON(http.StatusOK, gin.H{
			"message": "device deleted successfully",
		})
	}
}

// Update Device by ID method
func UpdateDeviceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameter value from the request
		deviceID := c.Query("id")

		//check if device exists
		if _, err := db.GetDeviceByID(deviceID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "device not found",
			})
			return
		}

		// bin user info from request
		var device models.Device
		if err := c.ShouldBindJSON(&device); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// update device in database
		db.UpdateDeviceByID(deviceID, device)

		c.JSON(http.StatusOK, gin.H{
			"message": "device updated successfully",
		})
	}
}
