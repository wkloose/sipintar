package services
import (
	"errors"
	"time"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/wkloose/tempproject.git/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserHearts(userID uuid.UUID) (int, int, error) {
	var heart models.HeartReward
	err := initializers.DB.First(&heart, "user_id = ?", userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		heart = models.HeartReward{
			ID:             uuid.New(),
			UserID:         userID,
			Hearts:         5,
			LastRegenTime:  time.Now(),
		}
		if err := initializers.DB.Create(&heart).Error; err != nil {
			return 0, 0, err
		}
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, 0, err
	}

	// Hitung jumlah nyawa yang harusnya diregenerasi
	now := time.Now()
	elapsed := now.Sub(heart.LastRegenTime)
	newHearts := int(elapsed.Minutes()) / 5

	if newHearts > 0 && heart.Hearts < 5 {
		total := heart.Hearts + newHearts
		if total > 5 {
			total = 5
		}
		heart.Hearts = total
		heart.LastRegenTime = now
		initializers.DB.Save(&heart)
	}

	// Hitung waktu ke nyawa berikutnya
	timeSince := time.Since(heart.LastRegenTime)
	elapsedSeconds := int(timeSince.Seconds())
	nextIn := 300 - (elapsedSeconds % 300)

	if heart.Hearts >= 5 {
		nextIn = 0
	}

	return heart.Hearts, nextIn, nil
}

func ClaimHeartAfterMaterial(userID uuid.UUID) error {
	var heart models.HeartReward
	err := initializers.DB.First(&heart, "user_id = ?", userID).Error
if errors.Is(err, gorm.ErrRecordNotFound) {
    // Inisialisasi heart baru
    heart = models.HeartReward{
        ID:            uuid.New(),
        UserID:        userID,
        Hearts:        0, // bisa juga 5, sesuai strategimu
        LastRegenTime: time.Now(),
    }
    if err := initializers.DB.Create(&heart).Error; err != nil {
        return err
    	}
	} else if err != nil {
    return err
	}


	if heart.Hearts >= 5 {
		return errors.New("maksimal nyawa sudah penuh")
	}

	heart.Hearts += 1
	if heart.Hearts > 5 {
		heart.Hearts = 5
	}
	initializers.DB.Save(&heart)
	return nil
}
