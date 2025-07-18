package models

import (
    "time"
    "github.com/google/uuid"
)

type Streak struct {
    ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID        uuid.UUID `gorm:"type:uuid;uniqueIndex"`
    CurrentStreak int
    LastUpdated   time.Time // tanggal terakhir belajar
}
