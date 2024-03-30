package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID   uint
	TotalAmount  float32
	IsComplete   bool
	OrderDetails []OrderDetails `gorm:"foreignKey:OrderID"`
	Payment      Payment        `gorm:"foreignKey:OrderID"`
}
