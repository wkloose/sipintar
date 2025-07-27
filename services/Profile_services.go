// services/profile_service.go
package services

import (
	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/wkloose/tempproject.git/models"
	"time"
)

type ProfileResponse struct {
	Username   string `json:"username"`
	Name       string `json:"name"`
	Description string `json:"description"`
	PhotoURL   string `json:"photo_url"`
	JoinedAt   time.Time `json:"joined_at"`

	Stats struct {
		LamaBelajar      int `json:"lama_belajar"`       
		SoalDikerjakan   int `json:"soal_dikerjakan"`    
		Penyelesaian     int `json:"penyelesaian"`      
		Streak           int `json:"streak"`
	} `json:"stats"`
}

func GetUserProfile(userID uuid.UUID) (*ProfileResponse, error) {
	var user models.User
	if err := initializers.DB.Preload("Profile").First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	var streak models.Streak
	_ = initializers.DB.First(&streak, "user_id = ?", userID).Error

	var scoreSessions []models.ScoreSession
	_ = initializers.DB.Where("user_id = ?", userID).Find(&scoreSessions).Error

	var progress []models.LearningProgress
	_ = initializers.DB.Where("user_id = ?", userID).Find(&progress).Error

	totalDuration := 0
	for _, session := range scoreSessions {
		totalDuration += session.Duration
	}

	totalSoal := 0
	for _, p := range progress {
		totalSoal += p.TotalQuestions
	}

	profile := &ProfileResponse{
		Username:   user.Username,
		Name:       user.Profile.Name,
		Description: user.Profile.Description,
		PhotoURL:   user.Profile.PhotoURL,
		JoinedAt:   user.CreatedAt,
	}

	profile.Stats.LamaBelajar = totalDuration
	profile.Stats.SoalDikerjakan = totalSoal
	profile.Stats.Penyelesaian = len(scoreSessions)
	profile.Stats.Streak = streak.CurrentStreak

	return profile, nil
}
type UpdateProfileInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PhotoURL    string `json:"photo_url"`
}

func UpdateUserProfile(userID uuid.UUID, input UpdateProfileInput) error {
	var profile models.UserProfile

	if err := initializers.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return err
	}

	profile.Name = input.Name
	profile.Description = input.Description
	profile.PhotoURL = input.PhotoURL

	return initializers.DB.Save(&profile).Error
}