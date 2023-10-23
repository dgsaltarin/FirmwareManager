package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SignUp function create a new user in dynamodb
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// bin user info from request
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dbUser models.User

		fmt.Println(user)

		// check if user already exists
		if err := db.DB.Where("username = ?", user.Username).First(&dbUser).Error; err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "user already exists",
			})
			return
		}

		// add uuid to user
		user.ID = uuid.New().String()

		// hashnig password
		err := user.GeneratePasswordHash()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error hashing password",
			})
			fmt.Println(err)
			return
		}

		// create new user in database
		db.DB.Create(&user)

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
