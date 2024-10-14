package entities

type BudgetCategory struct {
	BudgetID   uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
	Amount     uint
}
