// services/profile_service.go
package services

import (
	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/wkloose/tempproject.git/models"
	"time"
	"encoding/base64"
)

type ProfileResponse struct {
    Username    string    `json:"username"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    PhotoBase64 string    `json:"photo_base64"`
    JoinedAt    time.Time `json:"joined_at"`

    Stats struct {
        LamaBelajar    int `json:"lama_belajar"`
        SoalDikerjakan int `json:"soal_dikerjakan"`
        Penyelesaian   int `json:"penyelesaian"`
        Streak         int `json:"streak"`
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

    base64Img := ""
    if len(user.Profile.PhotoBlob) > 0 {
        base64Img = "data:image/png;base64," + base64.StdEncoding.EncodeToString(user.Profile.PhotoBlob)
    }

    profile := &ProfileResponse{
        Username:    user.Username,
        Name:        user.Profile.Name,
        Description: user.Profile.Description,
        PhotoBase64: base64Img,
        JoinedAt:    user.CreatedAt,
    }

    profile.Stats.LamaBelajar = totalDuration
    profile.Stats.SoalDikerjakan = totalSoal
    profile.Stats.Penyelesaian = len(scoreSessions)
    profile.Stats.Streak = streak.CurrentStreak

    return profile, nil
}


func UpdateUserProfile(userID uuid.UUID, name, description string, photoBlob []byte) error {
    var profile models.UserProfile

    if err := initializers.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
        return err
    }

    profile.Name = name
    profile.Description = description
    if len(photoBlob) > 0 {
        profile.PhotoBlob = photoBlob
    }

    return initializers.DB.Save(&profile).Error
}
