package models

import "gorm.io/gorm"

type IngredientStock struct {
	gorm.Model
	CreatedBy string
	ProductID uint
	Quatity   uint
	Movement  []StockMovement `gorm:"foreignKey:StockID"`
}
