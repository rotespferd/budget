package main

import (
	"fmt"
	"net/http"

	"github.com/rotespferd/budget/http/handler"
)

func main() {
	host := "localhost"
	port := "8080"

	http.HandleFunc("/budgets", handler.ListBudgetsHandler)

	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
