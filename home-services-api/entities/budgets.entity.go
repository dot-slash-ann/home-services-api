package entities

import (
	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	Name       string     `gorm:"not null"`
	Categories []Category `gorm:"many2many:budget_categories;foreignKey:ID;joinForeignKey:BudgetID;References:ID;joinReferences:CategoryID"`
}
