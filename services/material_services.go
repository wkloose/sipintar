package services

import (
    "github.com/google/uuid"
    "github.com/wkloose/tempproject.git/initializers"
    "github.com/wkloose/tempproject.git/models"
)

func GetAllMaterials() ([]models.Material, error) {
    var materials []models.Material
    err := initializers.DB.Preload("Contents").Find(&materials).Error
    return materials, err
}

func GetMaterialByID(id uuid.UUID) (models.Material, error) {
    var material models.Material
    err := initializers.DB.Preload("Contents").First(&material, "id = ?", id).Error
    return material, err
}
