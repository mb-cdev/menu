package databasemigrate

import (
	authmodel "menu/auth/model"
	database "menu/common/database"
	commonmodel "menu/common/model"
	restaurantmodel "menu/restaurant/model"
)

func MigrateDb() {
	database.DB.AutoMigrate(&authmodel.User{})
	database.DB.AutoMigrate(&commonmodel.Address{})
	database.DB.AutoMigrate(&commonmodel.Country{})
	database.DB.AutoMigrate(&restaurantmodel.Restaurant{})
	database.DB.AutoMigrate(&restaurantmodel.Menu{})
	database.DB.AutoMigrate(&restaurantmodel.Product{})
	database.DB.AutoMigrate(&commonmodel.Currency{})

	c := commonmodel.Country{Name: "Poland"}

	database.DB.Create(&c)

	cu := commonmodel.Currency{
		Name:          "Złoty",
		Symbol:        "zł",
		Exchange_rate: 1,
		Template:      "<amount> <symbol>",
	}

	database.DB.Create(&cu)

	cu2 := commonmodel.Currency{
		Name:          "Euro",
		Symbol:        "€",
		Exchange_rate: 0.2149,
		Template:      "<symbol><amount>",
	}

	database.DB.Create(&cu2)

	cu3 := commonmodel.Currency{
		Name:          "Dolar",
		Symbol:        "$",
		Exchange_rate: 0.2541,
		Template:      "<symbol><amount>",
	}

	database.DB.Create(&cu3)
}
