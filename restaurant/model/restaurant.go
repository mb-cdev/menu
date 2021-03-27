package model

import (
	authmodel "menu/auth/model"
	commonmodel "menu/common/model"

	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name      string
	AddressID uint
	Address   *commonmodel.Address
	UserID    uint
	User      *authmodel.User
	Menu      []*Menu
}
