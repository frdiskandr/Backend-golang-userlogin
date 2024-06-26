package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name           string `json:"name" validate:"required"`
    Email          string `json:"email" validate:"required,email"`
    Password       string `json:"password" validate:"required,min=6"`
    ProfilePicture string `json:"profile_picture"`
}
