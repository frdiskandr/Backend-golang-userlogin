package config

import (
	"api_btpn_logres/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    dsn := os.Getenv("DB_DSN")
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    DB.AutoMigrate(&models.User{})
}
