package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	FullName string
	Tel string
	Email string
	IsActive bool
	RoleID uint
}

type UserWithRole struct {
	User
	RoleName    string
}