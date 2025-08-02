package seed

import (
	"github.com/wkloose/tempproject.git/models"
)
	var LinearQuestions = []models.Question{
	{
		Type:       "multiple_choice",
		Question:    "Nilai dari $2x + 3 = 11$ adalah?",
		Options: []models.AnswerOption{
			{OptionText: "3", IsCorrect: false},
			{OptionText: "4", IsCorrect: true},
			{OptionText: "5", IsCorrect: false},
			{OptionText: "6", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika $5x = 25$, maka nilai $x =$ ...",
		Options: []models.AnswerOption{
			{OptionText: "4", IsCorrect: false},
			{OptionText: "5", IsCorrect: true},
			{OptionText: "6", IsCorrect: false},
			{OptionText: "10", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "$x - 7 = 0$ memiliki nilai $x =$",
		Options: []models.AnswerOption{
			{OptionText: "7", IsCorrect: true},
			{OptionText: "-7", IsCorrect: false},
			{OptionText: "0", IsCorrect: false},
			{OptionText: "1", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika $x = 3$, maka nilai $2x + 4 =$ ...",
		Options: []models.AnswerOption{
			{OptionText: "8", IsCorrect: false},
			{OptionText: "10", IsCorrect: true},
			{OptionText: "12", IsCorrect: false},
			{OptionText: "14", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika $y = 2x + 1$, maka $y = 7$ saat $x =$ ...",
		Options: []models.AnswerOption{
			{OptionText: "2", IsCorrect: false},
			{OptionText: "3", IsCorrect: true},
			{OptionText: "4", IsCorrect: false},
			{OptionText: "5", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Solusi dari $3x - 2 = 7$ adalah:",
		Options: []models.AnswerOption{
			{OptionText: "3", IsCorrect: true},
			{OptionText: "2", IsCorrect: false},
			{OptionText: "4", IsCorrect: false},
			{OptionText: "5", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika $2(x + 3) = 10$, maka $x =$ ...",
		Options: []models.AnswerOption{
			{OptionText: "1", IsCorrect: false},
			{OptionText: "2", IsCorrect: false},
			{OptionText: "3", IsCorrect: true},
			{OptionText: "4", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Tentukan titik potong sumbu-Y dari grafik $y = 3x - 5$",
		Options: []models.AnswerOption{
			{OptionText: "(0,3)", IsCorrect: false},
			{OptionText: "(0,5)", IsCorrect: false},
			{OptionText: "(0,-5)", IsCorrect: true},
			{OptionText: "(3,0)", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika dua garis sejajar, maka gradiennya ...",
		Options: []models.AnswerOption{
			{OptionText: "Sama", IsCorrect: true},
			{OptionText: "Berlawanan", IsCorrect: false},
			{OptionText: "Tidak tentu", IsCorrect: false},
			{OptionText: "Nol", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Garis $y = 2x + 1$ dan $y = -x + 4$ berpotongan di titik?",
		Options: []models.AnswerOption{
			{OptionText: "(1,3)", IsCorrect: true},
			{OptionText: "(2,2)", IsCorrect: false},
			{OptionText: "(0,1)", IsCorrect: false},
			{OptionText: "(3,1)", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "SPLDV:\n$$\n\\begin{cases}\nx + y = 5 \\\\\nx - y = 1\n\\end{cases}$$\nBerapakah nilai $x$?",
		Options: []models.AnswerOption{
			{OptionText: "2", IsCorrect: false},
			{OptionText: "3", IsCorrect: true},
			{OptionText: "4", IsCorrect: false},
			{OptionText: "5", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Koordinat yang tidak berada di garis $y = 2x$?",
		Options: []models.AnswerOption{
			{OptionText: "(1,2)", IsCorrect: false},
			{OptionText: "(2,4)", IsCorrect: false},
			{OptionText: "(3,6)", IsCorrect: false},
			{OptionText: "(2,5)", IsCorrect: true},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Nilai fungsi invers dari $f(x) = 2x - 1$ adalah:",
		Options: []models.AnswerOption{
			{OptionText: "$(x - 1)/2$", IsCorrect: true},
			{OptionText: "$(x + 1)/2$", IsCorrect: false},
			{OptionText: "$2x + 1$", IsCorrect: false},
			{OptionText: "$2x - 1$", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Grafik $y = -2x + 4$ memotong sumbu-X di titik ...",
		Options: []models.AnswerOption{
			{OptionText: "(0,4)", IsCorrect: false},
			{OptionText: "(2,0)", IsCorrect: true},
			{OptionText: "(-2,0)", IsCorrect: false},
			{OptionText: "(0,-2)", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "$y = mx + c$. Jika melalui titik (2,3) dan (4,7), maka $m =$ ...",
		Options: []models.AnswerOption{
			{OptionText: "1", IsCorrect: false},
			{OptionText: "2", IsCorrect: true},
			{OptionText: "3", IsCorrect: false},
			{OptionText: "4", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Solusi dari $x/2 = 6$ adalah:",
		Options: []models.AnswerOption{
			{OptionText: "12", IsCorrect: true},
			{OptionText: "6", IsCorrect: false},
			{OptionText: "3", IsCorrect: false},
			{OptionText: "2", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Hasil dari $4x - 2x = 12$ adalah $x =$ ...",
		Options: []models.AnswerOption{
			{OptionText: "3", IsCorrect: true},
			{OptionText: "6", IsCorrect: false},
			{OptionText: "8", IsCorrect: false},
			{OptionText: "10", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Jika $f(x) = x + 2$, maka $f(3) =$",
		Options: []models.AnswerOption{
			{OptionText: "5", IsCorrect: false},
			{OptionText: "6", IsCorrect: true},
			{OptionText: "7", IsCorrect: false},
			{OptionText: "4", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Garis $y = 3x - 9$ memiliki gradien ...",
		Options: []models.AnswerOption{
			{OptionText: "-3", IsCorrect: false},
			{OptionText: "3", IsCorrect: true},
			{OptionText: "0", IsCorrect: false},
			{OptionText: "9", IsCorrect: false},
		},
	},
	{
		Type:       "multiple_choice",
		Question:    "Persamaan $x + 2 = -5$ memiliki solusi $x =$ ...",
		Options: []models.AnswerOption{
			{OptionText: "-7", IsCorrect: true},
			{OptionText: "-3", IsCorrect: false},
			{OptionText: "3", IsCorrect: false},
			{OptionText: "7", IsCorrect: false},
		},
	},
	{
 		Type: "essay", 
 		Question: "Tentukan nilai $x$ dari persamaan $4x + 5 = 21$",
		Explanation: "4",
	},
 	{
		Type: "essay", 
		Question: "Selesaikan sistem persamaan berikut dengan metode eliminasi:\n\n$$\\begin{cases}2x + 3y = 12 \\\\ x + y = 5\\end{cases}$$", 
		Explanation: "x = 3, y = 1",
	},
 	{
		Type: "essay", 
		Question: "Buatlah persamaan garis lurus yang melalui titik (1, 2) dan (3, 6)",
		Explanation: "y = 2x",
	},
 	{
		Type: "essay", 
		Question: "Tentukan titik potong antara garis $y = 2x - 1$ dan $y = -x + 4$",
		Explanation: "(1, 1)",
	},
 	{
		Type: "essay",
	 	Question: "Jika diketahui $y = mx + c$, dan garis melalui titik (2, 3) dan (4, 5), berapakah nilai $m$ dan $c$?",
		Explanation: "m = 1, c = 1",
	},
	{
		Type: "essay", 
		Question: "Tentukan fungsi invers dari $f(x) = \\frac{3x + 2}{4}$",
		Explanation: "f⁻¹(x) = (4x - 2)/3",
	},
	{
		Type: "essay", 
		Question: "Ali dan Budi membeli 3 pensil dan 2 buku Rp12.000, dan 1 pensil serta 4 buku seharga Rp14.000. Tentukan harga masing-masing.",
		Explanation: "Pensil Rp2.000, Buku Rp3.000",
	},
	{
		Type: "essay", 
		Question: "Diberikan grafik garis lurus melalui titik (0,1) dan (4,5). Tentukan persamaan garis tersebut.",
		Explanation: "y = x + 1",
	},
 	{
		Type: "essay", 
		Question: "Jika $y = 2x + 3$, carilah titik potong sumbu X dari grafik tersebut.",
		Explanation: "Titik potong sumbu X adalah (-1.5, 0)",
	},
 	{
		Type: "essay", 
		Question: "Seseorang memiliki usia 3 kali lipat dari adiknya. Jumlah umur mereka 36 tahun. Buatlah model persamaan dan tentukan umur masing-masing.",
		Explanation: "Usia orang tua 27 tahun, adik 9 tahun",
	},
}