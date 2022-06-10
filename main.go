package main

import (
	"avg/avarage"
	"avg/types"
)

func main() {

	paym := []types.Payment{
		{
			ID:       1,
			Amount:   20,
			Category: "wallet",
		},
		{
			ID:       1,
			Amount:   30,
			Category: "Bonus",
		},
		{
			ID:       1,
			Amount:   40,
			Category: "Bonus",
		},
	}

	avarage.TotalInCategory(paym, "Bonus")
}
