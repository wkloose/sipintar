package initializers

import (
	"fmt"
	"os"

	"github.com/wkloose/tempproject.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println("failed to connect to database")
	}
}

func SyncDatabase() {
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.AnswerOption{},
		&models.HeartReward{},
		&models.LearningProgress{},
		&models.Material{},
		&models.MaterialContent{},
		&models.PasswordResetToken{},
		&models.Question{},
		&models.ScoreSession{},
		&models.Streak{},
		&models.UserProfile{},
	}

	for _, model := range modelsToMigrate {
		err := DB.AutoMigrate(model)
		if err != nil {
			fmt.Printf("Gagal migrate model: %T, error: %v\n", model, err)
		} else {
			fmt.Printf("Berhasil migrate model: %T (tabel mungkin sudah ada atau baru dibuat)\n", model)
		}
	}
}