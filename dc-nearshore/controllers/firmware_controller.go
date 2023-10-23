package controllers

import (
	"net/http"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/gin-gonic/gin"
)

// Create new Firmware in database
func CreateFirmware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// bin user info from request
		var firmware models.Firmware
		if err := c.ShouldBindJSON(&firmware); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// create new firmware in database
		firmware, err := db.CreateFirmware(firmware)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error creating firmware",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Firmware created successfully",
		})
	}
}

// Get Firmware by ID method
func GetFirmwareByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameter value from the request
		firmwareID := c.Query("id")
		var firmware models.Firmware

		// get firmware from database
		firmware, err := db.GetFirmwareByID(firmwareID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "firmware not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"firmware": firmware,
		})
	}
}

// Get All Firmwares method
func GetAllFirmwares() gin.HandlerFunc {
	return func(c *gin.Context) {
		var firmwares []models.Firmware

		// get users from database
		firmwares, err := db.GetAllFirmwares()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "firmwares not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"firmwares": firmwares,
		})
	}
}

// Delete Firmware by ID method
func DeleteFirmwareByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameter value from the request
		firmwareID := c.Query("id")

		//check if firmware exists
		if _, err := db.GetFirmwareByID(firmwareID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "firmware not found",
			})
			return
		}

		// delete firmware from database
		if err := db.DeleteFirmwareByID(firmwareID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "firmware not deleted",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "firmware deleted successfully",
		})
	}
}

// Update Firmware by ID method
func UpdateFirmwareByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameter value from the request
		firmwareID := c.Query("id")
		var firmware models.Firmware

		//check if firmware exists
		if _, err := db.GetFirmwareByID(firmwareID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "firmware not found",
			})
			return
		}

		// bin user info from request
		if err := c.ShouldBindJSON(&firmware); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// update firmware in database
		if err := db.UpdateFirmwareByID(firmwareID, firmware); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "firmware not updated",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "firmware updated successfully",
		})
	}
}
