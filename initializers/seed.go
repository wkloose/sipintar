package initializers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/models"
)

func SeedData() {
	var count int64
	DB.Model(&models.Material{}).Count(&count)
	if count > 0 {
		fmt.Println("Seed sudah pernah dijalankan, lewati.")
		return
	}

	material := models.Material{
		ID:          uuid.New(),
		Title:       "Bilangan Bulat",
		Description: "Materi tentang bilangan bulat untuk kelas 7",
	}
	DB.Create(&material)

	content1 := models.MaterialContent{
		ID:         uuid.New(),
		MaterialID: material.ID,
		Type:       "text",
		Text:       ptr("Bilangan bulat terdiri dari bilangan positif, negatif, dan nol."),
		Order:      1,
	}
	content2 := models.MaterialContent{
		ID:         uuid.New(),
		MaterialID: material.ID,
		Type:       "video",
		Link:       ptr("https://youtube.com/example-bilangan-bulat"),
		Order:      2,
	}
	DB.Create(&content1)
	DB.Create(&content2)

	question := models.Question{
		ID:         uuid.New(),
		MaterialID: material.ID,
		Type:       "multiple_choice",
		Question:   "Manakah di antara berikut ini yang merupakan bilangan bulat?",
		Explanation: "Bilangan bulat mencakup positif, negatif, dan nol.",
	}
	DB.Create(&question)

	options := []models.AnswerOption{
		{
			ID:         uuid.New(),
			QuestionID: question.ID,
			OptionText: "3.5",
			IsCorrect:  false,
		},
		{
			ID:         uuid.New(),
			QuestionID: question.ID,
			OptionText: "-2",
			IsCorrect:  true,
		},
		{
			ID:         uuid.New(),
			QuestionID: question.ID,
			OptionText: "0.8",
			IsCorrect:  false,
		},
		{
			ID:         uuid.New(),
			QuestionID: question.ID,
			OptionText: "√2",
			IsCorrect:  false,
		},
	}
	for _, opt := range options {
		DB.Create(&opt)
	}

	fmt.Println("✅ Seeding selesai.")
}

func ptr(s string) *string {
	return &s
}
