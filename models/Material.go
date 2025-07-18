package models

import (
    "github.com/google/uuid"
)

type Material struct {
    ID          uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Title       string            `gorm:"not null"`
    Description string
    Contents    []MaterialContent `gorm:"foreignKey:MaterialID"`
}
