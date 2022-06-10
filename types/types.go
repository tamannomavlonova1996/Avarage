package types

import "fmt"

type Money int64
type Category string

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

func (p Payment) GetAmountString() string {
	return fmt.Sprint(p.Amount)
}

func GetAmountString(payment *Payment) string {
	return fmt.Sprint(payment.Amount)
}
