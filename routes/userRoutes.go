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

	userGroup := r.Group("/users").Use(middlewares.VerifyToken(), middlewares.AccessPermission(userRepo))
	{
		userGroup.POST("/create", userController.CreateUser)
		userGroup.GET("/getOne/:id", userController.GetUser)
		// userGroup.PUT("/getOne/:id", userController.UpdateUser)
		// userGroup.DELETE("/getOne/:id", userController.DeleteUser)
		userGroup.GET("/getAlls", userController.GetAllUsers)
	}
}
