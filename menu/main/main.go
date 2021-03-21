package main

import (
	"fmt"
	"menu/common/database"
)

func init() {
	fmt.Println("Hello from main")
}

func main() {
	fmt.Println(database.DB)
}
