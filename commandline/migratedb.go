package main

import database "menu/common/databasemigrate"

func migrateDb() {
	database.MigrateDb()
}
