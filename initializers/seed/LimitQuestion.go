package seed

import (
	"github.com/wkloose/tempproject.git/models"
)
	var LimitQuestions = []models.Question{
	{
		Type:       "multiple_choice",
		Question:    "Nilai dari limₓ→2 (3x + 1) adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "5", IsCorrect: false},
			{OptionText: "6", IsCorrect: false},
			{OptionText: "7", IsCorrect: true},
			{OptionText: "8", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Hasil dari limₓ→1 (x² + 2x + 1) adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "2", IsCorrect: false},
			{OptionText: "4", IsCorrect: false},
			{OptionText: "6", IsCorrect: true},
			{OptionText: "9", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "limₓ→3 (x² - 9)/(x - 3) =",
		Options: []models.AnswerOption{
			{OptionText: "3", IsCorrect: false},
			{OptionText: "6", IsCorrect: false},
			{OptionText: "9", IsCorrect: true},
			{OptionText: "Tak tentu", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Nilai limₓ→0 (sin x / x) adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "1", IsCorrect: true},
			{OptionText: "∞", IsCorrect: false},
			{OptionText: "Tidak ada", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Nilai dari limₓ→∞ (3x + 2)/(x + 4) adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "1", IsCorrect: false},
			{OptionText: "2", IsCorrect: false},
			{OptionText: "3", IsCorrect: true},
			{OptionText: "∞", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "limₓ→0 (x / (x² + 1)) = ...",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: true},
			{OptionText: "1", IsCorrect: false},
			{OptionText: "Tidak ada", IsCorrect: false},
			{OptionText: "Tak hingga", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "limₓ→1 (x² - 1)/(x - 1)",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "2", IsCorrect: true},
			{OptionText: "1", IsCorrect: false},
			{OptionText: "∞", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika f(x) = {x + 2, x ≠ 3; k, x = 3} dan limₓ→3 f(x) = 5, maka nilai k = ...",
		Options: []models.AnswerOption{
			{OptionText: "2", IsCorrect: false},
			{OptionText: "3", IsCorrect: false},
			{OptionText: "5", IsCorrect: true},
			{OptionText: "6", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "limₓ→∞ (5x² + 1)/(2x² + 3) =",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "1", IsCorrect: false},
			{OptionText: "5/2", IsCorrect: true},
			{OptionText: "∞", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Nilai limₓ→0 (1 - cos x)/x² adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "1", IsCorrect: false},
			{OptionText: "1/2", IsCorrect: true},
			{OptionText: "∞", IsCorrect: false},
		},
	},

	{
		Type:       "multiple_choice",
		Question:    "Jika f(x) = (x² - 4)/(x - 2), maka limₓ→2 f(x) = ...",
		Options: []models.AnswerOption{
			{OptionText: "2", IsCorrect: false},
			{OptionText: "4", IsCorrect: true},
			{OptionText: "0", IsCorrect: false},
			{OptionText: "Tak tentu", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Nilai dari limₓ→-1 (x² - 1)/(x + 1) adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "-2", IsCorrect: false},
			{OptionText: "2", IsCorrect: true},
			{OptionText: "0", IsCorrect: false},
			{OptionText: "Tidak ada", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "limₓ→0 (tan x / x)",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "1", IsCorrect: true},
			{OptionText: "Tidak ada", IsCorrect: false},
			{OptionText: "∞", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika f(x) = (x² - 16)/(x - 4), maka limₓ→4 f(x) = ...",
		Options: []models.AnswerOption{
			{OptionText: "8", IsCorrect: true},
			{OptionText: "16", IsCorrect: false},
			{OptionText: "0", IsCorrect: false},
			{OptionText: "Tidak ada", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "limₓ→0 (ln(1 + x)/x)",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "1", IsCorrect: true},
			{OptionText: "e", IsCorrect: false},
			{OptionText: "∞", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "limₓ→∞ (2x³ + 1)/(5x³ + 4) =",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "1", IsCorrect: false},
			{OptionText: "2/5", IsCorrect: true},
			{OptionText: "∞", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika limit kiri dan kanan suatu fungsi di titik a tidak sama, maka ...",
		Options: []models.AnswerOption{
			{OptionText: "Limit ada", IsCorrect: false},
			{OptionText: "Limit tidak ada", IsCorrect: true},
			{OptionText: "Fungsi kontinu", IsCorrect: false},
			{OptionText: "Fungsi selalu naik", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "limₓ→2 (x² - 4)/(x - 2)",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "2", IsCorrect: false},
			{OptionText: "4", IsCorrect: true},
			{OptionText: "Tidak ada", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika f(x) = (x - 3)²/(x - 3), maka limₓ→3 f(x) =",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "3", IsCorrect: true},
			{OptionText: "6", IsCorrect: false},
			{OptionText: "∞", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Limit dari fungsi konstan f(x) = 7 untuk semua x adalah:",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: false},
			{OptionText: "7", IsCorrect: true},
			{OptionText: "Tidak ada", IsCorrect: false},
			{OptionText: "1", IsCorrect: false},
		},
	},

	{
		Type:       "essay",
		Question:    "Hitung limₓ→1 (x² - 1)/(x - 1)",
		Explanation:  "2",
	},
	{
		Type:       "essay",
		Question:    "Tentukan nilai a agar f(x) = (x² - a²)/(x - a) memiliki limit di x = a",
		Explanation:  "a ≠ 0",
	},
	{
		Type:       "essay",
		Question:    "Hitung limₓ→0 (tan x / x)",
		Explanation:  "1",
	},
	{
		Type:       "essay",
		Question:    "Diketahui f(x) = (x² - 9)/(x - 3). Hitung limₓ→3 f(x)",
		Explanation:  "6",
	},
	{
		Type:       "essay",
		Question:    "Hitung limₓ→0 (1 - cos x)/x²",
		Explanation:  "1/2",
	},
	{
		Type:       "essay",
		Question:    "Hitung limₓ→∞ (3x² + 1)/(2x² + 4)",
		Explanation:  "3/2",
	},
	{
		Type:       "essay",
		Question:    "Tentukan nilai limₓ→0 (ln(1 + x)/x)",
		Explanation:  "1",
	},
	{
		Type:       "essay",
		Question:    "Diketahui f(x) = (x² - 4x + 4)/(x - 2), hitung limₓ→2 f(x)",
		Explanation:  "0",
	},
	{
		Type:       "essay",
		Question:    "Hitung limₓ→0 (e^x - 1)/x",
		Explanation:  "1",
	},
	{
		Type:       "essay",
		Question:    "Jelaskan mengapa limₓ→a f(x) tidak ada jika limit kiri ≠ limit kanan",
		Explanation:  "Karena limit hanya ada jika limit kiri = limit kanan. Jika berbeda, maka limit tidak terdefinisi.",
	},
}

