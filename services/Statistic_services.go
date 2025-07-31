package services

import (
	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/wkloose/tempproject.git/models"
)

type ProgressStat struct {
	TotalQuestions int     `json:"total_questions"`
	Correct        int     `json:"correct"`
	Wrong          int     `json:"wrong"`
	CorrectPercent float64 `json:"correct_percent"`
	WrongPercent   float64 `json:"wrong_percent"`
}

func GetProgressStat(userID uuid.UUID) (*ProgressStat, error) {
	var progresses []models.LearningProgress
	if err := initializers.DB.Where("user_id = ?", userID).Find(&progresses).Error; err != nil {
		return nil, err
	}

	totalQuestions := 0
	totalCorrect := 0
	totalWrong := 0

	for _, p := range progresses {
		totalQuestions += p.TotalQuestions
		totalCorrect += p.CorrectAnswers
		totalWrong += p.WrongAnswers
	}

	stat := &ProgressStat{
		TotalQuestions: totalQuestions,
		Correct:        totalCorrect,
		Wrong:          totalWrong,
	}

	if totalQuestions > 0 {
		stat.CorrectPercent = float64(totalCorrect) / float64(totalQuestions) * 100
		stat.WrongPercent = float64(totalWrong) / float64(totalQuestions) * 100
	}

	return stat, nil
}
