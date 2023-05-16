package services

import (
	"testing"

	"github.com/raisa320/API/models"
)

type itemTest struct {
	item           models.Item
	resultExpected float64
}

var itemsTest = []itemTest{
	{models.Item{Price: 2, Quantity: 3}, 6},
	{models.Item{Price: 21, Quantity: 2}, 42},
	{models.Item{Price: 11, Quantity: 4}, 44},
	{models.Item{Price: 14, Quantity: 0}, 0},
}

func TestTotalPriceItem(t *testing.T) {
	for _, itemTest := range itemsTest {
		result := itemTest.item.GetTotalPrice()
		if result != itemTest.resultExpected {
			t.Errorf("Error en total price item expected %v value  %v", itemTest.resultExpected, result)
		}
	}
}

func TestIncrementViews(t *testing.T) {
	var increment = 0
	for i := 0; i < 5; i++ {
		go IncrementViews(&increment)
	}
}
