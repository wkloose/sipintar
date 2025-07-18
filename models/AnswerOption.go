package models

import (
    "github.com/google/uuid"
)

type AnswerOption struct {
    ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    QuestionID uuid.UUID
    OptionText string
    IsCorrect  bool
}
