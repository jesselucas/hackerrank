package main

import "testing"

func Test_absDiagonalDif(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected float64
	}{
		{[][]int{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}}, 0.0},
		{[][]int{[]int{11, 2, 4}, []int{4, 5, 6}, []int{10, 8, -12}}, 15.0},
		{[][]int{[]int{1}}, 0.0},
		{[][]int{[]int{1, 2}, []int{2, 1}}, 2.0},
	}

	for _, test := range tests {
		a := absDiagonalDif(test.matrix)
		if a != test.expected {
			t.Fatal("Received: ", a, "Expected: ", test.expected)
		}
	}
}
