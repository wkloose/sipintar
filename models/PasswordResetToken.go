package models

import (
    "time"
    "github.com/google/uuid"
)

type PasswordResetToken struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Email     string    `gorm:"not null"`
    Token     string    `gorm:"not null"`
    ExpiresAt time.Time
}
