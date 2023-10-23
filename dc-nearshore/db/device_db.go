package db

import "github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"

// GetDeviceByID method
func GetDeviceByID(deviceID string) (models.Device, error) {
	var device models.Device

	// get device from database
	if err := DB.Where("id = ?", deviceID).First(&device).Error; err != nil {
		return device, err
	}

	return device, nil
}

// DeleteDeviceByID method
func DeleteDeviceByID(deviceID string) error {
	var device models.Device

	// get device from database
	if err := DB.Where("id = ?", deviceID).Delete(&device).Error; err != nil {
		return err
	}

	return nil
}

// UpdateDeviceByID method
func UpdateDeviceByID(deviceID string, device models.Device) error {
	// update device in database
	if err := DB.Model(&device).Where("id = ?", deviceID).Updates(device).Error; err != nil {
		return err
	}

	return nil
}
