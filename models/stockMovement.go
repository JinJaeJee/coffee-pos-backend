package models

import "gorm.io/gorm"

type StockMovement struct {
	gorm.Model
	StockID       uint
	MovementType  MovementType
	QuatityChange uint
	NewQuatity    uint
	Reason        string
	RecordBy      string
}

type MovementType string

const (
	MovementAddition    MovementType = "addition"
	MovementSubtraction MovementType = "subtraction"
	MovementAdjustment  MovementType = "adjustment"
	MovementReturn      MovementType = "return"
)
