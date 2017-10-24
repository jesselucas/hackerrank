package main

import "testing"

func Test_mealCost(t *testing.T) {
	tests := []struct {
		meal     float64
		tip      float64
		tax      float64
		expected float64
	}{
		{10.0, 20.0, 6.0, 12.60},
	}

	for _, test := range tests {
		cost := mealCost(test.meal, test.tip, test.tax)
		if cost != test.expected {
			t.Fatal("Received: ", cost, "Expected: ", test.expected)
		}
	}
}
