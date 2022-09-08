package budget

import (
	"fmt"
	"net/http"

	"github.com/rotespferd/budget/common"
)

type ListBudgetStruct struct {
	Title   string
	Budgets []Budget
}

func ListBudgetsHandler(w http.ResponseWriter, r *http.Request) {
	db := common.GetDatabase()
	query := "SELECT * FROM budget"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	var budgets []Budget

	defer rows.Close()
	for rows.Next() {
		budget := NewBudgetFromDbRow(rows)
		budgets = append(budgets, budget)
	}

	data := ListBudgetStruct{
		Title:   "Budgets",
		Budgets: budgets,
	}

	w.Write([]byte(data.Title))
}
