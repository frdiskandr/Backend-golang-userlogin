package controllers

import (
	"api_btpn_logres/config"
	"api_btpn_logres/models"
	"api_btpn_logres/services"
	"api_btpn_logres/utils"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validate.Struct(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func LoginUser(c *gin.Context) {
    var input struct {
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validate.Struct(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := services.GetUserByEmail(input.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if !services.VerifyPassword(user.Password, input.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func UploadProfilePicture(c *gin.Context) {
    userId := c.MustGet("userId").(uint)
    file, err := c.FormFile("profile_picture")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
        return
    }

    filename := fmt.Sprintf("uploads/%d%s", userId, filepath.Ext(file.Filename))
    if err := c.SaveUploadedFile(file, filename); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := services.UpdateProfilePicture(userId, filename); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile picture uploaded successfully"})
}

func DeleteProfilePicture(c *gin.Context) {
    userId := c.MustGet("userId").(uint)

    if err := services.DeleteProfilePicture(userId); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile picture deleted successfully"})
}

func GetProfiles(c *gin.Context) {
    var users []models.User
    if err := config.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}