package stats

import (
	"github.com/bekhruzdilshod/bankdouble/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {

	totalSum := types.Money(0)

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			totalSum += payment.Amount
		}
		
	}

	return totalSum / types.Money(len(payments))

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
