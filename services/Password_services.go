package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/wkloose/tempproject.git/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"os"
)

func generateToken(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, n)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

func RequestPasswordReset(email string) (string, error) {
	var user models.User
	err := initializers.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("email tidak terdaftar")
		}
		return "", err 
	}

	expMinStr := os.Getenv("RESET_PASSWORD_EXPIRATION_MINUTES")
	expMin, err := strconv.Atoi(expMinStr)
	if err != nil || expMin <= 0 {
		expMin = 15 
	}

	token := generateToken(32)
	expiration := time.Now().Add(time.Duration(expMin) * time.Minute)

	resetToken := models.PasswordResetToken{
		ID:        uuid.New(),
		Email:     email,
		Token:     token,
		ExpiresAt: expiration,
	}

	err = initializers.DB.Create(&resetToken).Error
	if err != nil {
		fmt.Println("DEBUG Gagal insert resetToken:", err)
		return "", fmt.Errorf("Gagal membuat token reset")
	}

	fmt.Printf("Reset link: http://localhost:3000/reset-password?token=%s\n", token)

	return token, nil
}

func ResetPassword(token string, newPassword string) error {
	var resetToken models.PasswordResetToken
	err := initializers.DB.Where("token = ?", token).First(&resetToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("token tidak valid")
		}
		return err
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return fmt.Errorf("token sudah kadaluarsa")
	}

	var user models.User
	err = initializers.DB.Where("email = ?", resetToken.Email).First(&user).Error
	if err != nil {
		return fmt.Errorf("pengguna tidak ditemukan")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)
	err = initializers.DB.Save(&user).Error
	if err != nil {
		return err
	}

	initializers.DB.Delete(&resetToken)
	return nil
}
