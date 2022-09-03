package budget

import (
	"database/sql"
	"log"
)

type Budget struct {
	ID          int
	Name        string
	Description string
	Amount      int
}

func NewBudget(name, description string, amount, id int) Budget {
	return Budget{
		ID:          id,
		Name:        name,
		Description: description,
		Amount:      amount,
	}
}

func NewBudgetFromDbRow(row *sql.Rows) Budget {
	var id, amount int
	var name, description string
	err := row.Scan(&id, &name, &description, &amount)
	if err != nil {
		log.Fatal(err)
	}
	return NewBudget(name, description, amount, id)
}
