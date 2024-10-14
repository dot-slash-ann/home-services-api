package entities

import (
	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	Categories []Category `gorm:"many2many:budget_categories;foreignKey:ID;joinForeignKey:BudgetID;References:ID;joinReferences:CategoryID"`
}
