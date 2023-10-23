package controllers

import (
	"net/http"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// bin user info from request
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// add uuid to user
		user.ID = uuid.New().String()

		// create new user in database
		db.DB.Create(&user)

		c.JSON(200, gin.H{
			"message": "User created successfully",
		})
	}
}

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameter value from the request
		userID := c.Query("id")
		var user models.User

		// get users from database
		if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "user not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"users": user,
		})
	}
}

// get user by username	
func GetUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameter value from the request
		username := c.Query("username")
		var user models.User

		// get users from database
		if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "user not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"users": user,
		})
	}
}
