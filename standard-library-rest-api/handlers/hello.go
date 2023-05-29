package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements http.Handler
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Service is working...")
	d, err := ioutil.ReadAll(r.Body)
	// if we have error here
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s\n", d)
}
