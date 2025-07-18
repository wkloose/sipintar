package models

import (
    "time"
    "github.com/google/uuid"
)

type HeartReward struct {
    ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID        uuid.UUID `gorm:"type:uuid;uniqueIndex"`
    Hearts        int       // maksimum 5
    LastRegenTime time.Time // untuk logika regenerasi 10 menit
}
