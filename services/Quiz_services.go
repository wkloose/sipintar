package services

import (
    "math/rand"
    "time"

    "github.com/google/uuid"
    "github.com/wkloose/tempproject.git/initializers"
    "github.com/wkloose/tempproject.git/models"
	"fmt"
	"gorm.io/gorm"
)

func GetQuizQuestions(materialID uuid.UUID) ([]models.Question, error) {
    var allQuestions []models.Question

    err := initializers.DB.
        Preload("Options").
        Where("material_id = ?", materialID).
        Find(&allQuestions).Error

    if err != nil {
        return nil, err
    }

    // Acak soalnya
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(allQuestions), func(i, j int) {
        allQuestions[i], allQuestions[j] = allQuestions[j], allQuestions[i]
    })

    // Ambil maksimal 10 soal
    limit := 10
    if len(allQuestions) < 10 {
        limit = len(allQuestions)
    }

    return allQuestions[:limit], nil
}
type AnswerInput struct {
    QuestionID uuid.UUID `json:"question_id"`
    AnswerID   uuid.UUID `json:"answer_id"`
}

func SubmitQuiz(userID, materialID uuid.UUID, answers []AnswerInput) (int, int, error) {
    correct := 0
    wrong := 0

    for _, ans := range answers {
        var option models.AnswerOption
        err := initializers.DB.
            Where("id = ? AND question_id = ?", ans.AnswerID, ans.QuestionID).
            First(&option).Error
        if err != nil {
            wrong++
            continue
        }

        if option.IsCorrect {
            correct++
        } else {
            wrong++
        }
    }

    // Simpan ke ScoreSession
    session := models.ScoreSession{
        ID:         uuid.New(),
        UserID:     userID,
        MaterialID: materialID,
        Correct:    correct,
        Wrong:      wrong,
        Duration:   0, // opsional isi jika ada timer
        CreatedAt:  time.Now(),
    }

    if err := initializers.DB.Create(&session).Error; err != nil {
        return 0, 0, err
    }

    // Update LearningProgress
    var progress models.LearningProgress
    err := initializers.DB.
        Where("user_id = ? AND material_id = ?", userID, materialID).
        First(&progress).Error

    percent := float64(correct) / float64(correct+wrong) * 100

    if err != nil {
        progress = models.LearningProgress{
            ID:             uuid.New(),
            UserID:         userID,
            MaterialID:     materialID,
            TotalQuestions: correct + wrong,
            CorrectAnswers: correct,
            WrongAnswers:   wrong,
            Percent:        percent,
        }
        _ = initializers.DB.Create(&progress)
    } else {
        progress.TotalQuestions += correct + wrong
        progress.CorrectAnswers += correct
        progress.WrongAnswers += wrong
        progress.Percent = float64(progress.CorrectAnswers) / float64(progress.TotalQuestions) * 100
        _ = initializers.DB.Save(&progress)
    }

    // Update streak
    var streak models.Streak
	err = initializers.DB.Where("user_id = ?", userID).First(&streak).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			streak = models.Streak{
				ID:            uuid.New(),
				UserID:        userID,
				CurrentStreak: 1,
				LastUpdated:   time.Now(),
			}
			_ = initializers.DB.Create(&streak)
		} else {
			return correct, wrong, err
		}
	} else {
		if time.Since(streak.LastUpdated).Hours() < 24 {
			streak.CurrentStreak++
		} else {
			streak.CurrentStreak = 1
		}
		streak.LastUpdated = time.Now()
		_ = initializers.DB.Save(&streak)
	}

    // Kurangi nyawa
    _ = ReduceHeart(userID)

    return correct, wrong, nil
}

func ReduceHeart(userID uuid.UUID) error {
    var heart models.HeartReward
    err := initializers.DB.Where("user_id = ?", userID).First(&heart).Error
    if err != nil {
        return err
    }

    if heart.Hearts > 0 {
        heart.Hearts -= 1
    } else {
        return fmt.Errorf("nyawa tidak cukup")
    }

    return initializers.DB.Save(&heart).Error
}
