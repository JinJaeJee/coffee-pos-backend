package controllers

import (
	"coffee-pos-backend/models"
	"coffee-pos-backend/repositories"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userRepo *repositories.AuthRepository
}

// jwtKey := []byte.(Getenv("SECRET"))

func NewAuthController(userRepo *repositories.AuthRepository) *AuthController {
	return &AuthController{userRepo: userRepo}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var loginReq models.LoginRequest

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := ctrl.userRepo.FindByUsername(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	expirationTime := time.Now().Add(8 * time.Hour)
	claims := &models.Claims{
		Username: loginReq.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		log.Println("JWT_SECRET_KEY:", os.Getenv("JWT_SECRET_KEY"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Token: tokenString,
		User: models.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			FullName: user.FullName,
			Tel:      user.Tel,
			Email:    user.Email,
			IsActive: user.IsActive,
			RoleID:   user.RoleID,
		},
	})

}

func (ctrl *AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
