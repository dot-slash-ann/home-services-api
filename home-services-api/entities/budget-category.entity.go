package entities

import "gorm.io/gorm"

type BudgetCategory struct {
	gorm.Model
	BudgetID   uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
	Amount     uint
}
