package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/wkloose/tempproject.git/services"
	"github.com/wkloose/tempproject.git/models"
)

func GetQuizHandler(c *gin.Context) {
    materialIDStr := c.Param("material_id")
    materialID, err := uuid.Parse(materialIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Material ID tidak valid"})
        return
    }

    questions, err := services.GetQuizQuestions(materialID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil soal"})
        return
    }

    c.JSON(http.StatusOK, questions)
}
type SubmitQuizInput struct {
    Answers []services.AnswerInput `json:"answers" binding:"required"`
}

func SubmitQuizHandler(c *gin.Context) {
    userAny, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    user := userAny.(models.User)

    materialIDStr := c.Param("material_id")
    materialID, err := uuid.Parse(materialIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Material ID tidak valid"})
        return
    }

    var input SubmitQuizInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
        return
    }

    correct, wrong, err := services.SubmitQuiz(user.ID, materialID, input.Answers)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "correct": correct,
        "wrong":   wrong,
        "message": "Hasil kuis berhasil disimpan",
    })
}
