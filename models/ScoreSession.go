package models

import (
    "time"
    "github.com/google/uuid"
)

type ScoreSession struct {
    ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID     uuid.UUID `gorm:"type:uuid"`
    MaterialID uuid.UUID `gorm:"type:uuid"`
    Correct    int
    Wrong      int
    Duration   int       // dalam detik
    CreatedAt  time.Time
}
