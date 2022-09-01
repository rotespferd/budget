package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rotespferd/budget/http/handler"
)

func main() {
	host := "localhost"
	port := "8080"

	http.HandleFunc("/budgets", logging(handler.ListBudgetsHandler))

	log.Printf("Listening on %s:%s", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
