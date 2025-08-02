package models

import (
    "github.com/google/uuid"
)

type Question struct {
    ID          uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    MaterialID  uuid.UUID       `gorm:"type:uuid"` // relasi ke Material
    Type        string          // "multiple_choice" / "essay"
    Question    string
    Explanation string
    ImageURL    *string  
    Options     []AnswerOption  `gorm:"foreignKey:QuestionID"`
}
