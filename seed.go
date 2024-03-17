package main

import (
	"coffee-pos-backend/models"
	"log"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{RoleName: "super admin"},
        {RoleName: "admin"},
        {RoleName: "staff"},
        {RoleName: "vip member"},
        {RoleName: "member"},
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


