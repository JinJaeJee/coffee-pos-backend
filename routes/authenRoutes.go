package routes

import (
	"coffee-pos-backend/controllers"
	"coffee-pos-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewAuthRepository(db)
	authController := controllers.NewAuthController(userRepo)

	r.POST("/login", authController.Login)
	r.GET("/logout", authController.Logout)
}
