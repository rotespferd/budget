package budget

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
