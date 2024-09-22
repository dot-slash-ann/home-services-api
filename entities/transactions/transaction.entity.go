package TransactionsEntity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionOn time.Time `gorm:"type:DATE;not null;"`
	PostedOn      time.Time `gorm:"type:DATE;not null;"`
	Amount        int       `gorm:"not null;"`
	VendorId      int
	CategoryId    int
}
