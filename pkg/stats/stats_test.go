package stats

import (
	"reflect"
	"testing"

	"github.com/bekhruzdilshod/bankdouble/v2/pkg/types"
)

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("Result len != 0")
	}
}

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("Result len != 0")
	}

}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "mobile"},
		{ID: 3, Category: "auto"},
	}

	result := FilterByCategory(payments, "fun")

	if len(result) != 0 {
		t.Error("Result len != 0")
	}

}


func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "mobile"},
		{ID: 3, Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "mobile"},
	}

	result := FilterByCategory(payments, "mobile")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Invalid result! Expected: %v, actual: %v", expected, result)
	}

}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "mobile"},
		{ID: 3, Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Invalid result! Expected: %v, actual: %v", expected, result)
	}

}


func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{Amount: 1_000_000, Category: "auto"},
		{Amount: 1_000_000, Category: "auto"},
		{Amount: 500_000, Category: "fun"},
	}

	expected := map[types.Category]types.Money {
		"auto": 2_000_000,
		"fun": 500_000,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Invalid result! Expected: %v, Actual: %v", expected, result)
	}
}