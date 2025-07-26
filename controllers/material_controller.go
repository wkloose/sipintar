package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/wkloose/tempproject.git/services"
)

func GetAllMaterials(c *gin.Context) {
    materials, err := services.GetAllMaterials()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil materi"})
        return
    }
    c.JSON(http.StatusOK, materials)
}

func GetMaterialDetail(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
        return
    }

    material, err := services.GetMaterialByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Materi tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, material)
}
