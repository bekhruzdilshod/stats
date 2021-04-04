package stats

import (
	"github.com/bekhruzdilshod/bankdouble/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {

	totalSum := types.Money(0)
	length := 0 

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			totalSum += payment.Amount
			length = length + 1
		}
		
	}

	return totalSum / types.Money(length)

}

// TotalInCategory возвращает сумму совершенных платежей по переданной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	totalInCategory := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			totalInCategory = totalInCategory + payment.Amount
		}
	}

	return totalInCategory

}


// FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}


// CategoriesTotal возвращает сумму платежей по каждой категории 
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] = categories[payment.Category] + payment.Amount
	}

	return categories
}