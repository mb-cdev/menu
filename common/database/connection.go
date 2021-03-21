package database

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=127.0.0.1 dbname=menu user=postgres password=root"

var DB *gorm.DB

func init() {
	var once sync.Once

	once.Do(func() {
		var err error

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}
	})
}
