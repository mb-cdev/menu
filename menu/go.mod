module menu/menu

replace menu/common => ../common

replace menu/auth => ../auth

replace menu/restaurant => ../restaurant

replace menu/session => ../session

go 1.16

require menu/auth v0.0.0-00010101000000-000000000000
