module menu/auth

go 1.16

replace menu/common => ../common

replace menu/auth => ../auth

require (
	gorm.io/gorm v1.21.4
	menu/common v0.0.0-00010101000000-000000000000
)
