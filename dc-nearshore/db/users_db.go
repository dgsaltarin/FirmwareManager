package db

import (
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/google/uuid"
)

// Create new User in database
func CreateUser(user models.User) (models.User, error) {
	// add uuid to user
	user.ID = uuid.New().String()
	// create new user in database
	if err := DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Get User by ID
func GetUserByID(userID string) (models.User, error) {
	var user models.User

	// get user from database
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Get User by username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	// get user from database
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil
}
