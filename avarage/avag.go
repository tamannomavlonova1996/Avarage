package avarage

import (
	"avg/types"
	"fmt"
)

/*func Avg(payments []types.Payment) types.Money {
	var a types.Money
	var b types.Money
	for _, result := range payments {
		a += result.Amount
	}
	b = a / types.Money(len(payments))

	return b
}*/
func TotalInCategory(payments []types.Payment, category types.Category) {
	var a types.Money
	for _, result := range payments {
		if category == result.Category {
			a = a + result.Amount
		}

	}

	fmt.Println(a)
}
