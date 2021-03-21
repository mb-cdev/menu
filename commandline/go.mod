module menu/commandline

go 1.16

replace menu/common => ../common

replace menu/auth => ../auth

require (
	menu/auth v0.0.0-00010101000000-000000000000
	menu/common v0.0.0-00010101000000-000000000000
)
