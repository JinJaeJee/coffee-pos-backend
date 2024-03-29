package main

import (
	"coffee-pos-backend/models"
	"log"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{
			RoleName:    "super admin",
			Description: "Super Admin has all permissions.",
			Permission: models.RolePermission{
				UserView:   true,
				UserCreate: true,
				UserEdit:   true,
				MenuView:   true,
				MenuCreate: true,
				MenuEdit:   true,
				RoleView:   true,
				RoleEdit:   true,
			},
		},
		{
			RoleName:    "admin",
			Description: "Admin has limited permissions.",
			Permission: models.RolePermission{
				UserView:   true,
				UserCreate: true,
				UserEdit:   true,
				MenuView:   true,
				MenuCreate: true,
				MenuEdit:   false, // Example: Admin cannot edit menus
				RoleView:   true,
				RoleEdit:   false, // Example: Admin cannot edit roles
			},
		},
		{
			RoleName:    "staff",
			Description: "staff has mostly viewing permissions.",
			Permission: models.RolePermission{
				UserView:   false,
				UserCreate: false,
				UserEdit:   false,
				MenuView:   true,
				MenuCreate: false,
				MenuEdit:   false,
				RoleView:   false,
				RoleEdit:   false,
			},
		},
	}
	for _, role := range roles {
		err := db.Where(models.Role{RoleName: role.RoleName}).FirstOrCreate(&role).Error
		if err != nil {
			log.Printf("Role '%v' already exists or cannot be created: %v\n", role.RoleName, err)
		} else {
			log.Printf("Role '%v' seeded successfully.\n", role.RoleName)
		}
	}
}
