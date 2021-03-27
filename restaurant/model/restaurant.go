package model

import (
	authmodel "menu/auth/model"
	commonmodel "menu/common/model"
)

type Restaurant struct {
	Name    string
	Address *commonmodel.Address
	UserID  uint
	User    authmodel.User
}
