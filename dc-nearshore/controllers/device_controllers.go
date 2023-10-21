package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Device method
func CreateDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Create Device",
		})
	}
}
