package model

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	name         string
	Products     []Product
	RestaurantID uint
	valid_from   time.Time
	valid_to     time.Time
}
