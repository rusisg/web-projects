package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("func goodbye is working correctly")
	d, err := ioutil.ReadAll(r.Body)
	// if we have error here
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Goodbye %s\n", d)
}
