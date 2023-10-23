package db

import (
	"time"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/google/uuid"
)

// CreateFirmware method
func CreateFirmware(firmware models.Firmware) (models.Firmware, error) {
	// add infor to firmware
	firmware.ID = uuid.New().String()
	firmware.CreatedAt = time.Now()
	firmware.UpdatedAt = time.Now()

	// create new firmware in database
	if err := DB.Create(&firmware).Error; err != nil {
		return models.Firmware{}, err
	}

	return firmware, nil
}

// GetFirmwareByID method
func GetFirmwareByID(firmwareID string) (models.Firmware, error) {
	var firmware models.Firmware

	// get firmware from database
	if err := DB.Where("id = ?", firmwareID).First(&firmware).Error; err != nil {
		return firmware, err
	}

	return firmware, nil
}

// GetAllFirmwares method
func GetAllFirmwares() ([]models.Firmware, error) {
	var firmwares []models.Firmware

	// get firmwares from database
	if err := DB.Find(&firmwares).Error; err != nil {
		return firmwares, err
	}

	return firmwares, nil
}

// Delete Firmware by ID method
func DeleteFirmwareByID(firmwareID string) error {
	var firmware models.Firmware

	// get firmware from database
	if err := DB.Where("id = ?", firmwareID).Delete(&firmware).Error; err != nil {
		return err
	}

	return nil
}

// Update Firmware by ID method
func UpdateFirmwareByID(firmwareID string, firmware models.Firmware) error {
	// update firmware in database
	if err := DB.Model(&firmware).Where("id = ?", firmwareID).Updates(firmware).Error; err != nil {
		return err
	}

	return nil
}
