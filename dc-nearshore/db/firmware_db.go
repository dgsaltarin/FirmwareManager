package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
)

// CreateFirmware creates a new firmware record in the database.
func CreateFirmware(db *gorm.DB, firmware *models.Firmware) error {
	result := db.Create(firmware)
	return result.Error
}

// GetFirmwareByID retrieves a firmware record from the database by its ID.
func GetFirmwareByID(db *gorm.DB, id uuid.UUID) (*models.Firmware, error) {
	var firmware models.Firmware
	result := db.First(&firmware, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &firmware, nil
}

// UpdateFirmware updates an existing firmware record in the database.
func UpdateFirmware(db *gorm.DB, firmware *models.Firmware) error {
	result := db.Save(firmware)
	return result.Error
}

// DeleteFirmware deletes a firmware record from the database by its ID.
func DeleteFirmware(db *gorm.DB, id uuid.UUID) error {
	result := db.Delete(&models.Firmware{}, id)
	return result.Error
}
