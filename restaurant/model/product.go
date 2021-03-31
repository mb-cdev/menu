package model

import (
	"menu/common/model"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	MenuID uint
	Price  model.Price
	name   string
}
