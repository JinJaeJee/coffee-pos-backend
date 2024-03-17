package routes

import (
	controllers "coffee-pos-backend/controllers"
	"coffee-pos-backend/middlewares"

	"github.com/gin-gonic/gin"
)

type AuthenRoutes struct {
	app *gin.Engine
	authenCTRL *controllers.AuthController
	authenMDW *middlewares.AuthenticatorMDW
}

func NewAuthenRoute(app *gin.Engine) *AuthenRoutes {
	return &AuthenRoutes{
		app: app,
		authenCTRL: controllers.NewAuthController(),
		authenMDW: middlewares.NewAuthenticatorMDW(),
	}
}

func (r *AuthenRoutes) Setup(){
	AuthenRoutes:= r.app.Group("/auth")
	{
		AuthenRoutes.POST("/login", r.authenCTRL.Login)
		AuthenRoutes.POST("/logout", r.authenMDW.VerifyToken, r.authenCTRL.Logout)
	}
}