package models

import (
    "github.com/google/uuid"
)

type UserProfile struct {
    ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID         uuid.UUID `gorm:"type:uuid;uniqueIndex"` // relasi ke User
    Name           string
    Description    string
    PhotoURL       string
    TotalQuestions int
    TotalCorrect   int
    StreakID       uuid.UUID `gorm:"type:uuid"` // relasi ke Streak
    HeartRewardID  uuid.UUID `gorm:"type:uuid"` // relasi ke HeartReward
}
