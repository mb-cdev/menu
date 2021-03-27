module menu/common

go 1.16

replace menu/auth => ../auth

replace menu/common => ../common

replace menu/restaurant => ../restaurant

require (
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.6
	menu/auth v0.0.0-00010101000000-000000000000
	menu/restaurant v0.0.0-00010101000000-000000000000
)
