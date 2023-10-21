package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
)

// CreateDevice creates a new device record in the database.
func CreateDevice(db *gorm.DB, device *models.Device) error {
	result := db.Create(device)
	return result.Error
}

// GetDeviceByID retrieves a device record from the database by its ID.
func GetDeviceByID(db *gorm.DB, id uuid.UUID) (*models.Device, error) {
	var device models.Device
	result := db.First(&device, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &device, nil
}

// UpdateDevice updates an existing device record in the database.
func UpdateDevice(db *gorm.DB, device *models.Device) error {
	result := db.Save(device)
	return result.Error
}

// DeleteDevice deletes a device record from the database by its ID.
func DeleteDevice(db *gorm.DB, id uuid.UUID) error {
	result := db.Delete(&models.Device{}, id)
	return result.Error
}
