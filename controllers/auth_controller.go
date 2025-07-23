package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/wkloose/tempproject.git/services"
)

func Register(c *gin.Context) {
    var body struct {
        Username string `json:"username" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
        return
    }

    if err := services.RegisterUser(body.Username, body.Email, body.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil"})
}


func Login(c *gin.Context) {
    var body struct {
        Identifier string `json:"identifier"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, err := services.AuthenticateUser(body.Identifier, body.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.SetCookie("Authorization", token, 3600*24*7, "/", "", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
