package main

import "testing"

func Test_consecutiveOnes(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{13, 2},
		{3853, 4},
		{15, 4},
		{139, 2},
		{251, 5},
		{235, 3},
	}

	for i, test := range tests {
		c := consecutiveOnes(test.n)
		if c != test.expected {
			t.Fatalf("Test %d failed. Received: %d Expected: %d", i, c, test.expected)
		}
	}
}
