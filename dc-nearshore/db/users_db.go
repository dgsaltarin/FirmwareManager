package db

import "github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"

// Get User by username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	// get user from database
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil
}
