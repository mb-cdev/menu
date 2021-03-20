package common

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var once sync.Once

	once.Do(func() {
		var err error
		dsn := "host=127.0.0.1 user=postgres password=root dbname=menu post=5432"

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(fmt.Sprintln("Error while connecting database", err))
		}
	})
}
