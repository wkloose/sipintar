package models

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Email     string       `gorm:"unique;not null"`
    Password  string       `gorm:"not null"` // hashed
    CreatedAt time.Time
    UpdatedAt time.Time
    Profile   UserProfile  `gorm:"foreignKey:UserID"`
}
