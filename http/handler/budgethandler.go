package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type ListBudgetStruct struct {
	Title string
}

func ListBudgetsHandler(w http.ResponseWriter, r *http.Request) {
	data := ListBudgetStruct{
		Title: "Budgets",
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
