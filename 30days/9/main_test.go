package main

import "testing"

func Test_factorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{3, 6},
		{2, 2},
		{4, 24},
	}

	for _, test := range tests {
		f := factorial(test.n)
		if f != test.expected {
			t.Fatal("Received: ", f, "Expected: ", test.expected)
		}
	}
}
