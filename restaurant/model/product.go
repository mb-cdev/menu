package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	MenuID uint
	Price  Price `gorm:"type:numeric(14,2) not null default '0.00'"`
	name   string
}
