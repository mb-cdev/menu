module menu/auth

go 1.16

replace menu/common => ../common

replace menu/auth => ../auth

replace menu/restaurant => ../restaurant

replace menu/session => ../session

require (
	gorm.io/gorm v1.21.6
	menu/common v0.0.0-00010101000000-000000000000
	menu/session v0.0.0-00010101000000-000000000000
)
