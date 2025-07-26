package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/wkloose/tempproject.git/models"
    "github.com/wkloose/tempproject.git/services"
)

func GetStreakHandler(c *gin.Context) {
    userAny, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    user := userAny.(models.User)

    streak, err := services.GetStreakByUser(user.ID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Streak belum ada"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "current_streak": streak.CurrentStreak,
        "last_updated":   streak.LastUpdated.Format("2006-01-02"),
    })
}
