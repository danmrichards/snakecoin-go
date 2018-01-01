package models

import "fmt"

type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}

// Returns a string representation of the transaction.
func (t *Transaction) String() string {
	return fmt.Sprintf(
		"%s%s%d",
		t.From,
		t.To,
		t.Amount,
	)
}
