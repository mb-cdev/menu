module menu/session

go 1.16

replace menu/common => ../common

replace menu/auth => ../auth

replace menu/restaurant => ../restaurant

replace menu/session => ./

require (
	github.com/google/uuid v1.2.0
	menu/auth v0.0.0-00010101000000-000000000000
	menu/common v0.0.0-00010101000000-000000000000
)
