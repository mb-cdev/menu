package main

import (
	authmodel "menu/auth/model"
	"menu/common/database"
)

func migrateDb() {
	database.DB.AutoMigrate(&authmodel.User{})
}
