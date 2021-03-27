package model

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Name      string
	Line1     string
	Line2     string
	Postcode  string
	City      string
	CountryID uint
	Country   Country
	UserID    uint
}
