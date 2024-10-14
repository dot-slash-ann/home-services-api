package entities

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name string `gorm:"not null"`
}
