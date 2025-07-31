package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkloose/tempproject.git/services"
	"github.com/wkloose/tempproject.git/models"
)

func GetProgressStatHandler(c *gin.Context) {
	userAny, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user := userAny.(models.User)

	stat, err := services.GetProgressStat(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil statistik"})
		return
	}

	c.JSON(http.StatusOK, stat)
}
