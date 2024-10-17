package entities

import (
	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	Name string `gorm:"not null"`
}
