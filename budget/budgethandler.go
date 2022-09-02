package budget

import (
	"fmt"
	"html/template"
	"net/http"
)

type ListBudgetStruct struct {
	Title   string
	Budgets []Budget
}

func ListBudgetsHandler(w http.ResponseWriter, r *http.Request) {
	budgets := []Budget{
		NewBudget("Rent", "Rent for the apartment", 1000, 1),
		NewBudget("Food", "Food for the month", 500, 2),
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
