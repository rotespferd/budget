package handler

import (
	"fmt"
	"net/http"
)

func ListBudgetsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List budgets")
}
