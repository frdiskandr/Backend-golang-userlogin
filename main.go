package main

import (
	"api_btpn_logres/config"
	controllers "api_btpn_logres/controler"
	"api_btpn_logres/middleware"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Memuat variabel lingkungan dari file .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("gagal membaca file .env")
    }

    config.InitDB()

    r := gin.Default()
	r.Use(middleware.CORSMiddleware())

    r.POST("/register", controllers.RegisterUser)
    r.POST("/login", controllers.LoginUser)
    r.GET("/profiles", controllers.GetProfiles)

    protected := r.Group("/").Use(middleware.AuthMiddleware())
    {	
        protected.GET("/profile", controllers.GetProfile)
        protected.POST("/profile_picture", controllers.UploadProfilePicture)
        protected.PUT("/profile_picture", controllers.UpdateProfilePicture)
        protected.DELETE("/profile_picture", controllers.DeleteProfilePicture)
    }

    r.Run(":8080")
}
