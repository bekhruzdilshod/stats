package stats

import (
	"fmt"
	"github.com/bekhruzdilshod/bankdouble/v2/pkg/types"
)


func ExampleAvg() {

	payments := []types.Payment{
		{
			Amount: 100,
			Category: "Common",
		},
		{
			Amount: 100,
		},
		{
			Amount: 100,
		},
	}

	result := Avg(payments)
	fmt.Println(result)

	// Output: 100

}


func ExampleTotalInCategory() {

	payments := []types.Payment{
		{
			Category: "Продукты",
			Amount: 500,
		},
		{
			Category: "АЗС",
			Amount: 1000,
		},
		{
			Category: "Продукты",
			Amount: 300,
		},
	}

	result := TotalInCategory(payments, "Продукты")

	fmt.Println(result)

	// Output: 800


}