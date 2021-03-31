package model

import (
	"menu/common/model"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	MenuID      uint
	Price       model.Price
	Currency    model.Currency
	CurrencyID  uint
	Name        string `gorm:"type: VARCHAR(255) NOT NULL"`
	Description string `gorm:"type: VARCHAR(512) NOT NULL"`
}
