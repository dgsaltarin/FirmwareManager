package db

import "github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"

// GetFirmwareByID method
func GetFirmwareByID(firmwareID string) (models.Firmware, error) {
	var firmware models.Firmware

	// get firmware from database
	if err := DB.Where("id = ?", firmwareID).First(&firmware).Error; err != nil {
		return firmware, err
	}

	return firmware, nil
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
