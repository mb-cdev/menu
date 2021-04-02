package main

import (
	_ "menu/auth/backendregister"
	"net/http"
)

func main() {

	http.ListenAndServe(":6666", nil)
}
