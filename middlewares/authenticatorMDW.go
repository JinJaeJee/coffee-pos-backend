package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticatorMDW struct {
}

func NewAuthenticatorMDW() *AuthenticatorMDW {
	return &AuthenticatorMDW{}
}

func (mdw *AuthenticatorMDW) VerifyToken(c *gin.Context) {

	///// TO DO JWT TOKEN LATER
	token := c.GetHeader("Authorization")
	if token != "valid-token" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func (mdw *AuthenticatorMDW) AccessPermission(c *gin.Context) {
	///// todo permission logic later
	c.Next()


}