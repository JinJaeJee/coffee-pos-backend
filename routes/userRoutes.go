package routes

import (
	controllers "coffee-pos-backend/controllers/user"
	"coffee-pos-backend/middlewares"

	"github.com/gin-gonic/gin"
)


type UserRoutes struct {
	app *gin.Engine
	userCTRL *controllers.UserController
	userMDW *middlewares.AuthenticatorMDW

}

func NewUserRoute(app *gin.Engine) *UserRoutes {
	return &UserRoutes{
		app: app,
		userCTRL: controllers.NewUserController(),
		userMDW: middlewares.NewAuthenticatorMDW(),
	}
}


func (r *UserRoutes) Setup(){
	UserRoutes := r.app.Group("/users")
	{
		UserRoutes.POST("/create", r.userCTRL.CreateUser)

	}
}