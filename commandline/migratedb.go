package main

import "menu/common/database"

func migrateDb() {
	database.MigrateDb()
}
