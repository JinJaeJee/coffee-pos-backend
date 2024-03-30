package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID       uint
	Amount        float32
	PaymentMethod PaymentMethodType
}

type PaymentMethodType string

const (
	PaymentMethodPromtPay PaymentMethodType = "promtPay"
	PaymentMethodCredit   PaymentMethodType = "creditCard"
	PaymentMethodCash     PaymentMethodType = "cash"
)
