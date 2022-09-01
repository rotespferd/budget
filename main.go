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

	http.HandleFunc("/budgets", handler.ListBudgetsHandler)

	log.Printf("Listening on %s:%s", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
