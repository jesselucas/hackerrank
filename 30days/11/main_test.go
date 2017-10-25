package main

import "testing"

func Test_hourglassMaxSum(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected int
	}{
		{[][]int{
			[]int{1, 1, 1, 0, 0, 0},
			[]int{0, 1, 0, 0, 0, 0},
			[]int{1, 1, 1, 0, 0, 0},
			[]int{0, 0, 2, 4, 4, 0},
			[]int{0, 0, 0, 2, 0, 0},
			[]int{0, 0, 1, 2, 4, 0},
		}, 19},
		{[][]int{
			[]int{-1, -1, 0, -9, -2, -2},
			[]int{-2, -1, -6, -8, -2, -5},
			[]int{-1, -1, -1, -2, -3, -4},
			[]int{-1, -9, -2, -4, -4, -5},
			[]int{-7, -3, -3, -2, -9, -9},
			[]int{-1, -3, -1, -2, -4, -5},
		}, -6},
	}

	for _, test := range tests {
		sum := hourglassMaxSum(test.matrix)

		if sum != test.expected {
			t.Fatal("Received: ", sum, "Expected: ", test.expected)
		}
	}
}
