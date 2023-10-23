package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/db"
	"github.com/dgsaltarin/FirmwareManager/dc-nearshore/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user from request
		var userRequest *models.User
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userDb models.User

		// get users from database
		if err := db.DB.Where("username = ?", userRequest.Username).First(&userDb).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "invalid username or password",
			})
			return
		}

		// check if password is correct
		if userDb.CheckPassword(userRequest.Password) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": userDb.Username,
				"exp": time.Now().Add(time.Minute * 10).Unix(),
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
			c.JSON(http.StatusOK, gin.H{
				"user": userDb.Username,
			})

		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Invalid email or password",
			})
		}

	}
}
