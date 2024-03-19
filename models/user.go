package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	FullName     string
	Tel          string
	Email        string
	IsActive     bool
	RoleID       uint
}

// UserInfo DTO to be used in API responses
type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	RoleID   uint   `json:"role_id"`
}

type UserWithRole struct {
	UserInfo
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
