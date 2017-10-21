package main

import "testing"

func Test_sumInts(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 10, 11}, 31},
	}

	for _, test := range tests {
		sum := sumInts(test.a)
		if sum != test.expected {
			t.Fatal("Received: ", sum, "Expected: ", test.expected)
		}
	}
}
