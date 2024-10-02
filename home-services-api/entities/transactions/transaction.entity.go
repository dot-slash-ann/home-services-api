package transactions

import (
	"time"

	CategoriesEntity "github.com/dot-slash-ann/home-services-api/entities/categories"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionOn time.Time `gorm:"type:DATE;not null;"`
	PostedOn      time.Time `gorm:"type:DATE;not null;"`
	Amount        uint      `gorm:"not null;"`
	VendorId      uint
	CategoryID    uint
	Category      CategoriesEntity.Category `gorm:"foreignKey:CategoryID;references:ID"`
}
