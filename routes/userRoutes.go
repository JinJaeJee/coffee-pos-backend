package routes

import (
	"coffee-pos-backend/controllers"
	"coffee-pos-backend/middlewares"
	"coffee-pos-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	userGroup := r.Group("/users").Use(middlewares.VerifyToken())
	{
		userGroup.POST("/create", userController.CreateUser)
		userGroup.GET("/:id", userController.GetUser)
		// userGroup.PUT("/:id", userController.UpdateUser)
		// userGroup.DELETE("/:id", userController.DeleteUser)
		userGroup.GET("/getAlls", userController.GetAllUsers)
	}
}
