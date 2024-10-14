package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name    string   `gorm:"not null"`
	Budgets []Budget `gorm:"many2many:budget_categories"`
}
