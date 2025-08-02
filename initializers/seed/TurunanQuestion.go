package seed

import (
	"github.com/wkloose/tempproject.git/models"
)

var Turunanquestions = []models.Question{
	{
		Type:       "multiple_choice",
		Question: "Turunan dari f(x) = x^2 adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "2x", IsCorrect: true},
			{OptionText: "x", IsCorrect: false},
			{OptionText: "x^3", IsCorrect: false},
			{OptionText: "2x^2", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Jika y = 3x^3 + 2x, maka dy/dx = ...",
		Options: []models.AnswerOption{
			{OptionText: "6x + 2", IsCorrect: false},
			{OptionText: "9x^2 + 2", IsCorrect: true},
			{OptionText: "6x^2", IsCorrect: false},
			{OptionText: "3x^2", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Turunan dari y = √x adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "1/(2x)", IsCorrect: false},
			{OptionText: "1/(2√x)", IsCorrect: true},
			{OptionText: "1/x^2", IsCorrect: false},
			{OptionText: "1/x", IsCorrect: false},
		},
	},
	{

		Type:       "multiple_choice",
		Question: "Jika f(x) = sin x, maka f'(x) = ...",
		Options: []models.AnswerOption{
			{OptionText: "cos x", IsCorrect: true},
			{OptionText: "-sin x", IsCorrect: false},
			{OptionText: "-cos x", IsCorrect: false},
			{OptionText: "sin x", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "d/dx(5x^4) = ...",
		Options: []models.AnswerOption{
			{OptionText: "5x^3", IsCorrect: false},
			{OptionText: "20x^3", IsCorrect: true},
			{OptionText: "4x^5", IsCorrect: false},
			{OptionText: "15x^2", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Turunan dari y = e^x adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "e^{x+1}", IsCorrect: false},
			{OptionText: "x · e^x", IsCorrect: false},
			{OptionText: "e^x", IsCorrect: true},
			{OptionText: "e^{2x}", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Turunan dari f(x) = ln x adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "1/x", IsCorrect: true},
			{OptionText: "ln x", IsCorrect: false},
			{OptionText: "x ln x", IsCorrect: false},
			{OptionText: "x", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Jika y = 3x^3 - 4x^2 + x, maka y' = ...",
		Options: []models.AnswerOption{
			{OptionText: "9x^2 - 8x + 1", IsCorrect: true},
			{OptionText: "6x - 4 + 1", IsCorrect: false},
			{OptionText: "9x - 8x + 1", IsCorrect: false},
			{OptionText: "3x^2 - 8x", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Fungsi f(x) = x^2 + 2x memiliki turunan ...",
		Options: []models.AnswerOption{
			{OptionText: "2x + 2", IsCorrect: true},
			{OptionText: "2x", IsCorrect: false},
			{OptionText: "2x - 2", IsCorrect: false},
			{OptionText: "x + 2", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "d/dx(x^3 - 2x^2 + 4x - 7) adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "3x^2 + 2x + 4", IsCorrect: false},
			{OptionText: "3x^2 - 4x + 4", IsCorrect: true},
			{OptionText: "3x^2 - 4x", IsCorrect: false},
			{OptionText: "2x - 4", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Turunan dari f(x) = cos x adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "sin x", IsCorrect: false},
			{OptionText: "-sin x", IsCorrect: true},
			{OptionText: "-cos x", IsCorrect: false},
			{OptionText: "cos x", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Jika y = x^2 + 5, maka dy/dx = ...",
		Options: []models.AnswerOption{
			{OptionText: "2x + 5", IsCorrect: false},
			{OptionText: "2x", IsCorrect: true},
			{OptionText: "x^2", IsCorrect: false},
			{OptionText: "5", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Turunan dari fungsi konstan f(x) = 7 adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "0", IsCorrect: true},
			{OptionText: "1", IsCorrect: false},
			{OptionText: "7", IsCorrect: false},
			{OptionText: "Tidak terdefinisi", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "d/dx(x^3 + x^2 + x + 1) = ...",
		Options: []models.AnswerOption{
			{OptionText: "3x^2 + 2x + 1", IsCorrect: true},
			{OptionText: "3x + 2 + 1", IsCorrect: false},
			{OptionText: "x^3 + x^2 + 1", IsCorrect: false},
			{OptionText: "3x^2 + x", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Jika f(x) = tan x, maka f'(x) = ...",
		Options: []models.AnswerOption{
			{OptionText: "sec^2 x", IsCorrect: true},
			{OptionText: "tan x", IsCorrect: false},
			{OptionText: "sin x", IsCorrect: false},
			{OptionText: "sec x", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "f(x) = 1/x, maka turunan f'(x) = ...",
		Options: []models.AnswerOption{
			{OptionText: "-1/x^2", IsCorrect: true},
			{OptionText: "1/x^2", IsCorrect: false},
			{OptionText: "ln x", IsCorrect: false},
			{OptionText: "1", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Turunan dari f(x) = x^{-2} adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "-2x^{-3}", IsCorrect: true},
			{OptionText: "2x^{-3}", IsCorrect: false},
			{OptionText: "-x^{-2}", IsCorrect: false},
			{OptionText: "-2x^{-1}", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Jika s(t) = 4t^2 + 2t, maka kecepatan saat t = 1 adalah ...",
		Options: []models.AnswerOption{
			{OptionText: "6", IsCorrect: false},
			{OptionText: "8", IsCorrect: true},
			{OptionText: "10", IsCorrect: false},
			{OptionText: "12", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Aturan turunan dari f(x) = u(x) · v(x) dikenal sebagai ...",
		Options: []models.AnswerOption{
			{OptionText: "Aturan Rantai", IsCorrect: false},
			{OptionText: "Aturan Perkalian", IsCorrect: true},
			{OptionText: "Aturan Pangkat", IsCorrect: false},
			{OptionText: "Aturan Limit", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question: "Jika y = (2x + 1)^3, maka dy/dx = ...",
		Options: []models.AnswerOption{
			{OptionText: "3(2x + 1)^2 · 2", IsCorrect: true},
			{OptionText: "3(2x + 1)^2", IsCorrect: false},
			{OptionText: "6x + 1", IsCorrect: false},
			{OptionText: "6x^2", IsCorrect: false},
		},
	},

	// ✍️ Essay Questions
	{
		Type:       "essay",
		Question: "Hitung turunan dari y = x^3 + 2x",
		Explanation:  "y' = 3x^2 + 2",
	},
	{
		Type:       "essay",
		Question: "Hitung d/dx(5x^2 - 7x + 1)",
		Explanation:  "10x - 7",
	},
	{
		Type:       "essay",
		Question: "Jika f(x) = (x^2 + 1)(x - 3), hitung turunan pertama f'(x)",
		Explanation:  "f'(x) = 2x(x - 3) + x^2 + 1",
	},
	{
		Type:       "essay",
		Question: "Hitung turunan dari y = (3x + 1)/(x - 2)",
		Explanation:  "y' = -7/(x - 2)^2",
	},
	{
		Type:       "essay",
		Question: "Jika f(x) = √(3x + 2), hitung f'(x)",
		Explanation:  "f'(x) = 3 / (2√(3x + 2))",
	},
	{
		Type:       "essay",
		Question: "Turunkan y = ln(2x + 1)",
		Explanation:  "y' = 2 / (2x + 1)",
	},
	{
		Type:       "essay",
		Question: "Sebuah benda bergerak menurut s(t) = t^3 - 6t^2 + 9t. Berapa kecepatan saat t = 2?",
		Explanation:  "s'(2) = -3",
	},
	{
		Type:       "essay",
		Question: "Hitung turunan dari y = tan(3x)",
		Explanation:  "y' = 3 sec^2(3x)",
	},
	{
		Type:       "essay",
		Question: "Turunan dari y = (x^2 + 1)^5 adalah?",
		Explanation:  "y' = 10x(x^2 + 1)^4",
	},
	{
		Type:       "essay",
		Question: "Tentukan nilai x agar turunan dari f(x) = x^2 - 4x + 3 sama dengan 0",
		Explanation:  "x = 2",
	},
}
