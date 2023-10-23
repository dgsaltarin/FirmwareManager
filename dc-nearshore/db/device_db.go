package db

import (
	"time"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/google/uuid"
)

// CreateDevice method
func CreateDevice(device models.Device) (models.Device, error) {
	// add infor to device
	device.ID = uuid.New().String()
	device.CreatedAt = time.Now()
	device.UpdatedAt = time.Now()

	// create new device in database
	if err := DB.Create(&device).Error; err != nil {
		return models.Device{}, err
	}

	return device, nil
}

// GetDeviceByID method
func GetDeviceByID(deviceID string) (models.Device, error) {
	var device models.Device

	// get device from database
	if err := DB.Where("id = ?", deviceID).First(&device).Error; err != nil {
		return device, err
	}

	return device, nil
}

// GetAllDevices method
func GetAllDevices() ([]models.Device, error) {
	var devices []models.Device

	// get devices from database
	if err := DB.Find(&devices).Error; err != nil {
		return devices, err
	}

	return devices, nil
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
