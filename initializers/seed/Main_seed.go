package seed

import (
	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/wkloose/tempproject.git/models"
	"log"
)

func seedMaterial(title, description, imageURL, videoURL, textContent string, questions []models.Question) {
	var existingMaterial models.Material
	if err := initializers.DB.Where("title = ?", title).First(&existingMaterial).Error; err == nil {
	initializers.DB.Where("material_id = ?", existingMaterial.ID).Delete(&models.MaterialContent{})
	var oldQuestions []models.Question
	initializers.DB.Where("material_id = ?", existingMaterial.ID).Find(&oldQuestions)
	for _, q := range oldQuestions {
		initializers.DB.Where("question_id = ?", q.ID).Delete(&models.AnswerOption{})
	}
	initializers.DB.Where("material_id = ?", existingMaterial.ID).Delete(&models.Question{})
	initializers.DB.Delete(&existingMaterial)
	}

	material := models.Material{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
	}
	initializers.DB.Create(&material)

	contents := []models.MaterialContent{
		{MaterialID: material.ID, Type: "text", Text: &textContent, Order: 1},
		{MaterialID: material.ID, Type: "image", Link: &imageURL, Order: 2},
		{MaterialID: material.ID, Type: "video", Link: &videoURL, Order: 3},
	}
	for _, content := range contents {
		initializers.DB.Create(&content)
	}

	for _, q := range questions {
	q.ID = uuid.New()
	q.MaterialID = material.ID

	opts := q.Options    
	q.Options = nil        

	if err := initializers.DB.Create(&q).Error; err != nil {
		log.Printf("Gagal menyimpan soal: %v\n", err)
		continue
	}

	if q.Type == "multiple_choice" {
		for _, opt := range opts {
			opt.ID = uuid.New()
			opt.QuestionID = q.ID
			if err := initializers.DB.Create(&opt).Error; err != nil {
				log.Printf("Gagal menyimpan opsi jawaban: %v\n", err)
			}
		}
	}
}
}
func SeedLinear() {
	seedMaterial(
		"Persamaan Linear",
		"Materi persamaan linear satu variabel.",
		"https://raw.githubusercontent.com/wkloose/sipintar/main/asset/persamaanlinear1.png",
		"https://youtu.be/Z-RhrP4eOJ8?si=SObNJGvs_JjG_t1K",
		"Persamaan linear adalah persamaan aljabar yang setiap sukunya mengandung paling tinggi satu variabel dengan pangkat satu.Bentuk umum persamaan linear dengan dua variabel adalah ax + by = c, di mana a, b, dan c adalah konstanta, dan x serta y adalah variabel.Grafik dari persamaan linear akan selalu membentuk garis lurus pada bidang kartesius",
		LinearQuestions,
	)
}

func SeedLimit() {
	seedMaterial(
		"Limit",
		"Materi limit fungsi aljabar dan trigonometri.",
		"https://raw.githubusercontent.com/wkloose/sipintar/main/asset/limit.png",
		"https://youtu.be/HEbMs2zYoDw?si=GeYV1N8LGrJ45zGA",
		"Limit dalam matematika pada dasarnya adalah nilai yang dituju oleh suatu fungsi ketika inputnya mendekati suatu bilangan tertentu.Jadi, kita tidak melihat nilai fungsi tepat di titik itu, melainkan nilai yang didekati oleh fungsi saat inputnya semakin dekat dan dekat ke titik tersebut.Sebagai contoh, jika kita memiliki fungsi dan ingin mengetahui limitnya saat x mendekati 0, kita akan melihat ke mana nilai fungsi tersebut bergerak saat x semakin mendekati 0, baik dari sisi positif maupun negatif. Perlu diingat bahwa nilai limit bisa saja berbeda dengan nilai eksak fungsi di titik tersebut.Berikut adalah gambar yang dapat membantu Anda memahami konsep limit:",
		LimitQuestions,
	)
}

func SeedTurunan() {
	seedMaterial(
		"Turunan",
		"Materi turunan fungsi aljabar dan trigonometri.",
		"https://raw.githubusercontent.com/wkloose/sipintar/main/asset/Turunan.png",
		"https://youtu.be/UtTvAVOOD6U?si=sUjA5I-6sffXYA2_",
		"Turunan dalam matematika mengukur seberapa cepat suatu fungsi berubah terhadap perubahan pada variabelnya. Secara geometris, turunan di suatu titik pada grafik fungsi merepresentasikan kemiringan garis singgung kurva di titik tersebut.Proses menemukan turunan suatu fungsi disebut diferensiasi. Turunan sangat penting dalam berbagai bidang ilmu, termasuk fisika (untuk menghitung kecepatan dan percepatan), ekonomi (untuk menganalisis perubahan marginal), dan teknik (untuk optimasi desain).Berikut adalah gambar yang mengilustrasikan turunan sebagai kemiringan garis singgung pada kurva:",
		Turunanquestions,
	)
}
func SeedIntegral() {
	seedMaterial(
		"Integral",
		"Materi integral tak tentu dan tentu.",
		"https://raw.githubusercontent.com/wkloose/sipintar/main/asset/integral.png",
		"https://youtu.be/E86ckq8yLUU?si=GX7V239m9-ASpoqK",
		"Integral adalah kebalikan dari operasi turunan, atau bisa juga disebut sebagai anti-turunan.Jika turunan mencari laju perubahan suatu fungsi, maka integral bertujuan untuk mengembalikan fungsi yang sudah diturunkan ke bentuk aslinya.Ada dua jenis integral:Integral Tak Tentu: Ini adalah proses untuk menemukan fungsi asli dari fungsi yang telah diturunkan. Karena konstanta akan hilang saat diturunkan, hasil dari integral tak tentu selalu menyertakan '+ C' untuk mewakili konstanta yang tidak diketahui.Integral Tentu: Integral ini memiliki batas atas dan bawah, dan hasilnya adalah nilai numerik yang spesifik. Integral tentu sering digunakan untuk menghitung luas daerah di bawah kurva.Berikut adalah gambar yang mengilustrasikan konsep integral sebagai luas di bawah kurva:",
		IntegralQuestions,
	)
}