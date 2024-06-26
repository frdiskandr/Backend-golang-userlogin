package services

import (
	"api_btpn_logres/config"
	"api_btpn_logres/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return config.DB.Create(user).Error
}

func GetUserByEmail(email string) (models.User, error) {
    var user models.User
    err := config.DB.Where("email = ?", email).First(&user).Error
    return user, err
}

func VerifyPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

func UpdateProfilePicture(userId uint, filePath string) error {
    return config.DB.Model(&models.User{}).Where("id = ?", userId).Update("profile_picture", filePath).Error
}

func DeleteProfilePicture(userId uint) error {
    return config.DB.Model(&models.User{}).Where("id = ?", userId).Update("profile_picture", "").Error
}
