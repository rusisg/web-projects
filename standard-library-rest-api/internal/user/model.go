package main

import (
	"fmt"
	"net/http"
)

func Something(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
func main() {
	http.HandleFunc("/", Something)

	http.ListenAndServe(":9090", nil)
}
