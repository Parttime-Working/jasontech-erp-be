package controllers

import (
	"erp/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var credentials models.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if credentials.Username == "admin" && credentials.Password == "password" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}
