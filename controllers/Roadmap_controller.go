package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/wkloose/tempproject.git/models"
    "github.com/wkloose/tempproject.git/services"
)

func GetRoadmap(c *gin.Context) {
    userAny, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    user := userAny.(models.User)

    roadmap, err := services.GetRoadmapForUser(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil roadmap"})
        return
    }

    c.JSON(http.StatusOK, roadmap)
}

func GetTotalProgressHandler(c *gin.Context) {
    userAny, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    user := userAny.(models.User)

    progress, err := services.GetTotalProgress(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung progress"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "total_progress": progress,
    })
}

