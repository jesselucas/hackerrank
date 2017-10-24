package main

import "testing"

func Test_sprintReverse(t *testing.T) {
	tests := []struct {
		a        []int
		expected string
	}{
		{[]int{1, 2, 3}, "3 2 1"},
		{[]int{7, 10, 8, 100}, "100 8 10 7"},
	}

	for _, test := range tests {
		out := sprintReverse(test.a)
		if out != test.expected {
			t.Fatal("Received: ", out, "Expected: ", test.expected)
		}
	}
}
