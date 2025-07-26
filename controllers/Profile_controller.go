package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkloose/tempproject.git/services"
	"github.com/wkloose/tempproject.git/models"
)

func GetProfile(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := userInterface.(models.User)
	profile, err := services.GetUserProfile(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data profil"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
func UpdateProfile(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := userInterface.(models.User)

	var input services.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	err := services.UpdateUserProfile(user.ID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan perubahan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perubahan berhasil dilakukan"})
}