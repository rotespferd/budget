package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/rotespferd/budget/budget"
	"github.com/rotespferd/budget/common"
	"github.com/rotespferd/budget/common/http/middleware"
)

func main() {
	// assign value of env HOST to variable host, if not set use localhost
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	// assign value of env PORT to variable port, if not set use 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// initialize templates
	common.SetTemplateConfig("./template/", "./template/_layout")
	common.LoadTemplates()

	http.HandleFunc("/budgets", middleware.Chain(budget.ListBudgetsHandler, middleware.Method("GET"), middleware.Logging()))

	log.Printf("Listening on %s:%s", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}