package model

import (
	"menu/common/database"

	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	Name          string
	Symbol        string
	Template      string
	Exchange_rate float32 `gorm:"type:numeric(5,4) not null"`
}

func GetMainCurrency() *Currency {
	c := &Currency{}
	database.DB.Where("exchange_rate = ?", 1.0000).First(c)
	return c
}
