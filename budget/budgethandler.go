package budget

import (
	"fmt"
	"html/template"
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

	// load template from file /template/budget/index.html
	tmpl, err := template.ParseFiles("template/budget/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	// render template with data
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
