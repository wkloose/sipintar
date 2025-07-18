package models

import (
    "github.com/google/uuid"
)

type LearningProgress struct {
    ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID         uuid.UUID
    MaterialID     uuid.UUID
    TotalQuestions int
    CorrectAnswers int
    WrongAnswers   int
    Percent        float64 // otomatis dihitung
}
