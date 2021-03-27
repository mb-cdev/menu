module menu/restaurant

go 1.16

replace menu/auth => ../auth

replace menu/common => ../common

require (
	menu/auth v0.0.0-00010101000000-000000000000
	menu/common v0.0.0-00010101000000-000000000000
)
