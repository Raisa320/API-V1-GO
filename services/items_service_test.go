package services

import "testing"

type itemTest struct {
	quantity       int
	price          float64
	resultExpected float64
}

var itemsTest = []itemTest{
	{2, 12, 24},
	{0, 12, 0},
	{7, 3, 21},
}

func TestTotalPriceItem(t *testing.T) {
	for _, item := range itemsTest {
		result := TotalPriceItem(item.quantity, item.price)
		if result != item.resultExpected {
			t.Errorf("Error en total price item expected %v value  %v", item.resultExpected, result)
		}
	}
}
