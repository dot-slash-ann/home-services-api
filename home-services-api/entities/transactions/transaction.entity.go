package transactions

import (
	"time"

	"github.com/dot-slash-ann/home-services-api/entities/categories"
	"github.com/dot-slash-ann/home-services-api/entities/tags"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionOn time.Time `gorm:"type:DATE;not null;"`
	PostedOn      time.Time `gorm:"type:DATE;not null;"`
	Amount        uint      `gorm:"not null;"`
	VendorId      uint
	CategoryID    uint
	Category      categories.Category `gorm:"foreignKey:CategoryID;references:ID"`
	Tags          []tags.Tag          `gorm:"many2many:transaction_tags;foreignKey:ID;joinForeignKey:TransactionID;References:ID;joinReferences:TagID"`
}
