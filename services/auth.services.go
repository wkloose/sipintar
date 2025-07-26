package services

import (
    "errors"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "github.com/wkloose/tempproject.git/initializers"
    "github.com/wkloose/tempproject.git/models"
	"gorm.io/gorm"
)

func RegisterUser(username, email, password string) error {
     var existingUser models.User

    err := initializers.DB.
        Where("username = ? OR email = ?", username, email).
        First(&existingUser).Error

    if err == nil {
        return errors.New("email atau username sudah digunakan")
    }

    if !errors.Is(err, gorm.ErrRecordNotFound) {
        return err
    }
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := models.User{
        ID:        uuid.New(),
        Username:  username,
        Email:     email,
        Password:  string(hashed),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    result := initializers.DB.Create(&user)
    if result.Error != nil {
        return result.Error
    }
    profile := models.UserProfile{
        ID:       uuid.New(),
        UserID:   user.ID,
        Name:     "",         // bisa default
        PhotoURL: "",         // bisa dikosongkan dulu
    }
     _ = initializers.DB.Create(&profile)

    return nil
}


func AuthenticateUser(identifier, password string) (string, error) {
    var user models.User
    result := initializers.DB.First(&user, "email = ? OR username = ?", identifier,identifier)
    if result.Error != nil {
		return "", errors.New("invalid email/username or password")
	}

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid email/username or password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID.String(),
        "exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
