package main

import (
	_ "menu/auth/login"
	_ "menu/auth/register"
	"net/http"
)

func main() {
	http.ListenAndServe(":6666", nil)
}
