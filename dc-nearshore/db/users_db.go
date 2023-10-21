package db

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"gorm.io/gorm"
)

const userstable = "sharedbites-users"

// DatabaseInterface interface for the data type of db field
type DatabaseInterface interface {
	Upsert(table string, model interface{}) error
	GetItem(table string, model interface{}) (*dynamodb.GetItemOutput, error)
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

// UserDB storage structure
type UsersDb struct {
	db DatabaseInterface
}

// CreateUser creates a new user record in the database.
func CreateUser(db *gorm.DB, user *models.User) error {
	result := db.Create(user)
	return result.Error
}

// GetUserByID retrieves a user record from the database by its ID.
func GetUserByID(db *gorm.DB, id string) (*models.User, error) {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser updates an existing user record in the database.
func UpdateUser(db *gorm.DB, user *models.User) error {
	result := db.Save(user)
	return result.Error
}

// DeleteUser deletes a user record from the database by its ID.
func DeleteUser(db *gorm.DB, id string) error {
	result := db.Delete(&models.User{}, id)
	return result.Error
}
