package handler

import (
	"fmt"
	"log"
	"net/http"
)

func ListBudgetsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ListBudgetsHandler called")

	fmt.Fprintf(w, "List budgets")
}
