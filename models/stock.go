package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	CreatedBy string
	ProductID uint
	Quatity   uint
	Movement  []StockMovement `gorm:"foreignKey:StockID"`
}
