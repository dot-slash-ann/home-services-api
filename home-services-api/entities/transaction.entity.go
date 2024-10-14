package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionOn   time.Time `gorm:"type:DATE;not null;"`
	PostedOn        time.Time `gorm:"type:DATE;not null;"`
	Amount          uint      `gorm:"not null;"`
	TransactionType string    `gorm:"not null;"`
	CategoryID      uint
	Category        Category `gorm:"foreignKey:CategoryID;references:ID"`
	VendorID        uint
	Vendor          Vendor `gorm:"foreignKey:VendorID;references:ID"`
	Tags            []Tag  `gorm:"many2many:transaction_tags;foreignKey:ID;joinForeignKey:TransactionID;References:ID;joinReferences:TagID"`
}
