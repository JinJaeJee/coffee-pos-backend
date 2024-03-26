package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName    string `gorm:"unique;not null"`
	Description string
	Permission  RolePermission `gorm:"foreignKey:RoleID"`
}
