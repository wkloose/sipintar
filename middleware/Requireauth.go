package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/wkloose/tempproject.git/models"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	sub, ok := claims["sub"].(string)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(sub)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	initializers.DB.First(&user, "id = ?", userID)

	if user.ID == uuid.Nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	c.Set("userID", user.ID)
	c.Next()
}
