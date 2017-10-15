package main

import "testing"

func Test_howManyBlown(t *testing.T) {
	tests := []struct {
		candles  []int
		expected int
	}{
		{[]int{3, 1, 2, 3}, 2},
		{[]int{4, 1, 2, 3}, 1},
		{[]int{1, 1, 1, 1}, 4},
	}

	for _, test := range tests {
		blown := howManyBlown(test.candles)
		if blown != test.expected {
			t.Fatal("Received: ", blown, "Expected: ", test.expected)
		}
	}
}
