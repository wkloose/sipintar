package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkloose/tempproject.git/services"
	"github.com/wkloose/tempproject.git/models"
	"io"
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

    name := c.PostForm("name")
    description := c.PostForm("description")

    file, err := c.FormFile("photo")
    var photoBlob []byte
    if err == nil {
        f, err := file.Open()
        defer f.Close()
        if err == nil {
            photoBlob, _ = io.ReadAll(f)
        }
    }

    err = services.UpdateUserProfile(user.ID, name, description, photoBlob)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan perubahan"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Perubahan berhasil dilakukan"})
}
