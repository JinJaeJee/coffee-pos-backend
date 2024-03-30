package models

import "gorm.io/gorm"

type OrderDetails struct {
	gorm.Model
	OrderID        uint
	ProductID      uint
	Quantity       uint
	DiscountNumber float32
	DiscountType   DiscountType
	PriceAtOrder   float32
}

type DiscountType string

const (
	DiscountPercent DiscountType = "percent"
	DiscountFixed   DiscountType = "fixed"
)
