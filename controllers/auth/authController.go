package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ctrl *AuthController) Login(c *gin.Context) {

	//// TODO add logic later 
	c.JSON(http.StatusOK, gin.H{"message": "Login seccessful"})
}

func (ctrl *AuthController) Logout(c *gin.Context) {

	//// TODO add logic later 

	
	c.JSON(http.StatusOK, gin.H{"message": "Logout !!!!"})
}