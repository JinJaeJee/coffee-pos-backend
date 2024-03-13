package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	RoleId uint 
	Role Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}