package seed

import (
	"github.com/wkloose/tempproject.git/models"
)
	var IntegralQuestions = []models.Question{
	{
		Type:       "multiple_choice",
		Question:    "∫ x² dx =",
		Options: []models.AnswerOption{
			{OptionText: "x³/2 + C", IsCorrect: false},
			{OptionText: "x³/3 + C", IsCorrect: true},
			{OptionText: "2x + C", IsCorrect: false},
			{OptionText: "3x² + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ (2x + 3) dx =",
		Options: []models.AnswerOption{
			{OptionText: "x² + 3x + C", IsCorrect: true},
			{OptionText: "2x² + 3x + C", IsCorrect: false},
			{OptionText: "x² + x + C", IsCorrect: false},
			{OptionText: "2x + 3 + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ dx =",
		Options: []models.AnswerOption{
			{OptionText: "x + C", IsCorrect: true},
			{OptionText: "1 + C", IsCorrect: false},
			{OptionText: "0 + C", IsCorrect: false},
			{OptionText: "1/x + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ 3 dx =",
		Options: []models.AnswerOption{
			{OptionText: "3x + C", IsCorrect: true},
			{OptionText: "x + 3 + C", IsCorrect: false},
			{OptionText: "x³ + C", IsCorrect: false},
			{OptionText: "3/x + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ cos x dx =",
		Options: []models.AnswerOption{
			{OptionText: "sin x + C", IsCorrect: true},
			{OptionText: "-sin x + C", IsCorrect: false},
			{OptionText: "cos x + C", IsCorrect: false},
			{OptionText: "-cos x + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ 1/x dx =",
		Options: []models.AnswerOption{
			{OptionText: "ln |x| + C", IsCorrect: true},
			{OptionText: "1/x² + C", IsCorrect: false},
			{OptionText: "x + C", IsCorrect: false},
			{OptionText: "e^x + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ e^x dx =",
		Options: []models.AnswerOption{
			{OptionText: "x · e^x + C", IsCorrect: false},
			{OptionText: "e^x + C", IsCorrect: true},
			{OptionText: "e^(x+1) + C", IsCorrect: false},
			{OptionText: "ln x + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ x dx dari 0 sampai 2 adalah:",
		Options: []models.AnswerOption{
			{OptionText: "2", IsCorrect: true},
			{OptionText: "4", IsCorrect: false},
			{OptionText: "2", IsCorrect: false},
			{OptionText: "2", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Integral tentu ∫₀¹ (x² + x) dx hasilnya ...",
		Options: []models.AnswerOption{
			{OptionText: "3/2", IsCorrect: false},
			{OptionText: "5/6", IsCorrect: true},
			{OptionText: "1", IsCorrect: false},
			{OptionText: "1/2", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ 4x³ dx =",
		Options: []models.AnswerOption{
			{OptionText: "x⁴ + C", IsCorrect: true},
			{OptionText: "x³ + C", IsCorrect: false},
			{OptionText: "x⁵ + C", IsCorrect: false},
			{OptionText: "x⁴ + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ x⁻¹ dx adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "ln |x| + C", IsCorrect: true},
			{OptionText: "x²/2 + C", IsCorrect: false},
			{OptionText: "1/x + C", IsCorrect: false},
			{OptionText: "x + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika f(x) = x², maka luas daerah dari x = 0 ke x = 2 adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "8/3", IsCorrect: true},
			{OptionText: "4", IsCorrect: false},
			{OptionText: "2", IsCorrect: false},
			{OptionText: "3", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ (6x + 5) dx =",
		Options: []models.AnswerOption{
			{OptionText: "3x² + 5x + C", IsCorrect: true},
			{OptionText: "6x² + 5x + C", IsCorrect: false},
			{OptionText: "x² + 5 + C", IsCorrect: false},
			{OptionText: "x³ + 5x + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika v(t) = 4t, maka jarak dari t = 0 ke t = 3 adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "6", IsCorrect: false},
			{OptionText: "8", IsCorrect: false},
			{OptionText: "12", IsCorrect: false},
			{OptionText: "18", IsCorrect: true},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ x³ dx =",
		Options: []models.AnswerOption{
			{OptionText: "x⁴/4 + C", IsCorrect: true},
			{OptionText: "4x³/3 + C", IsCorrect: false},
			{OptionText: "x⁴ + C", IsCorrect: false},
			{OptionText: "3x⁴/4 + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika d/dx F(x) = f(x), maka F(x) adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "Integral tak tentu dari f(x)", IsCorrect: true},
			{OptionText: "Turunan dari f(x)", IsCorrect: false},
			{OptionText: "Turunan dari F(x)", IsCorrect: false},
			{OptionText: "Fungsi eksponen", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫ sec² x dx =",
		Options: []models.AnswerOption{
			{OptionText: "sec x + C", IsCorrect: false},
			{OptionText: "tan x + C", IsCorrect: true},
			{OptionText: "cot x + C", IsCorrect: false},
			{OptionText: "sin x + C", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫₁³ 2x dx =",
		Options: []models.AnswerOption{
			{OptionText: "4", IsCorrect: false},
			{OptionText: "6", IsCorrect: false},
			{OptionText: "8", IsCorrect: true},
			{OptionText: "8", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Integral dari fungsi konstan f(x) = 7 adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "7x + C", IsCorrect: true},
			{OptionText: "x + C", IsCorrect: false},
			{OptionText: "ln x + C", IsCorrect: false},
			{OptionText: "0", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "∫₀² x² dx =",
		Options: []models.AnswerOption{
			{OptionText: "4/3", IsCorrect: true},
			{OptionText: "2/3", IsCorrect: false},
			{OptionText: "2", IsCorrect: false},
			{OptionText: "8/3", IsCorrect: false},
		},
	},

	{
		Type:       "essay",
		Question:    "Hitung ∫ x dx",
		Explanation:  "½x² + C",
	},
	{
		Type:       "essay",
		Question:    "Hitung ∫ (2x² + 3x + 1) dx",
		Explanation:  "2/3x³ + 3/2x² + x + C",
	},
	{
		Type:       "essay",
		Question:    "Tentukan nilai ∫₁² x dx",
		Explanation:  "3/2",
	},
	{
		Type:       "essay",
		Question:    "Hitung luas daerah di bawah kurva f(x) = x dari x = 0 ke x = 4",
		Explanation:  "8",
	},
	{
		Type:       "essay",
		Question:    "Hitung ∫₀¹ (3x² + 2x) dx",
		Explanation:  "2",
	},
	{
		Type:       "essay",
		Question:    "Jika v(t) = t² + 2, tentukan jarak tempuh dari t = 0 ke t = 2",
		Explanation:  "20/3",
	},
	{
		Type:       "essay",
		Question:    "Hitung ∫ (x² - 4x + 4) dx",
		Explanation:  "⅓x³ - 2x² + 4x + C",
	},
	{
		Type:       "essay",
		Question:    "Hitung integral tentu ∫₁³ x² dx",
		Explanation:  "26/3",
	},
	{
		Type:       "essay",
		Question:    "Hitung ∫ sin x dx",
		Explanation:  "-cos x + C",
	},
	{
		Type:       "essay",
		Question:    "Jika f'(x) = 4x - 2, tentukan f(x)",
		Explanation:  "2x² - 2x + C",
	},
}

