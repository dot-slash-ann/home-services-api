package vendors

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}
