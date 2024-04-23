package main

import (
	"coffee-pos-backend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
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
	superAdminRole := models.Role{}
	db.Where("role_name = ?", "super admin").First(&superAdminRole)

	if superAdminRole.ID != 0 {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("hashed546###paSsword"), 14)
		if err != nil {
			log.Printf("Error hashing password: %v\n", err)
			return
		}

		superAdmin := models.User{
			Username:     "superadmin",
			PasswordHash: string(hashedPassword),
			FullName:     "Super Admin",
			Email:        "superadmin@coffee.com",
			IsActive:     true,
			RoleID:       1,
		}
		err = db.Create(&superAdmin).Error
		if err != nil {
			log.Printf("Super admin user creation failed: %v\n", err)
		} else {
			log.Println("Super admin user seeded successfully.")
		}
	} else {
		log.Println("Super admin role not found. Super admin user cannot be seeded.")
	}
}
