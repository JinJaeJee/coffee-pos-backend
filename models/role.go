package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string `gorm:"unique;not null"`
	User []User `gorm:"foreignKey:RoleID"`
}