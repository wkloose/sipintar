package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/wkloose/tempproject.git/models"
	"gorm.io/gorm"
)

func UpdateStreak(userID uuid.UUID) error {
	var streak models.Streak
	today := time.Now().Truncate(24 * time.Hour)

	err := initializers.DB.Where("user_id = ?", userID).First(&streak).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Belum ada streak → buat baru
		streak = models.Streak{
			ID:            uuid.New(),
			UserID:        userID,
			CurrentStreak: 1,
			LastUpdated:   today,
		}
		return initializers.DB.Create(&streak).Error
	} else if err != nil {
		return err
	}

	last := streak.LastUpdated.Truncate(24 * time.Hour)

	switch {
	case today.Equal(last):
		// Sudah update hari ini
		return nil
	case today.Sub(last) == 24*time.Hour:
		// Hari berikutnya → lanjutkan streak
		streak.CurrentStreak += 1
	default:
		// Lewat lebih dari 1 hari → reset
		streak.CurrentStreak = 1
	}

	streak.LastUpdated = today
	return initializers.DB.Save(&streak).Error
}
func GetStreakByUser(userID uuid.UUID) (models.Streak, error) {
	var streak models.Streak
	err := initializers.DB.Where("user_id = ?", userID).First(&streak).Error
	return streak, err
}
