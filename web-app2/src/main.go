package main

import (
	"net/http"
)

func main() {
	file := http.FileServer(http.Dir("assets/"))
	http.Handle("/", http.StripPrefix("/", file))

	http.ListenAndServe(":8080", nil)
}
