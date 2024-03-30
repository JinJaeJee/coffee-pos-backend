package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CreatedBy     string `gorm:"not null"`
	Fullname      string `gorm:"not null"`
	Address       string
	Email         string
	PhoneNumber   string
	LoyaltyPoints int
	Order         []Order `gorm:"foreignKey:CustomerID"`
}
