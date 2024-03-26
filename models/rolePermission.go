package models

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleID     uint
	UserView   bool
	UserCreate bool
	UserEdit   bool
	MenuView   bool
	MenuCreate bool
	MenuEdit   bool
	RoleView   bool
	RoleEdit   bool
}
