package models

import (
    "github.com/google/uuid"
)

type UserProfile struct {
    ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID         uuid.UUID `gorm:"type:uuid;uniqueIndex"` // relasi ke User
    Name           string
    Description    string
    PhotoBlob      []byte // Ganti dari PhotoURL
    TotalQuestions int
    TotalCorrect   int
    StreakID       uuid.UUID `gorm:"type:uuid"`
    HeartRewardID  uuid.UUID `gorm:"type:uuid"`
}

