package controllers

import (
	"net/http"
	"time"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		// add infor to firmware
		firmware.ID = uuid.New().String()
		firmware.CreatedAt = time.Now()
		firmware.UpdatedAt = time.Now()

		// create new firmware in database
		db.DB.Create(&firmware)

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

		// get users from database
		if err := db.DB.Where("id = ?", firmwareID).First(&firmware).Error; err != nil {
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
		if err := db.DB.Find(&firmwares).Error; err != nil {
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
