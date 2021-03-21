package main

import (
	"flag"
	"fmt"
)

var command string

func main() {
	flag.StringVar(&command, "c", "", "Execute command")
	flag.Parse()

	switch command {
	case "migrate":
		fmt.Println("Migrating database")
		migrateDb()
	default:
		fmt.Println("Choose your operation")
	}
}
