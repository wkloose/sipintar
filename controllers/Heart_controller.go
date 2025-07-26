package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/services"
)

func GetHeartCount(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	hearts, _, err := services.GetUserHearts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil nyawa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hearts": hearts})
}

func GetHeartStatus(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	hearts, nextIn, err := services.GetUserHearts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil status nyawa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hearts":        hearts,
		"next_heart_in": nextIn,
	})
}

func ClaimHeart(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	err := services.ClaimHeartAfterMaterial(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nyawa berhasil diklaim"})
}
