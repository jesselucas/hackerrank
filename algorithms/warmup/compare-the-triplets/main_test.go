package main

import "testing"

func Test_compareScores(t *testing.T) {
	tests := []struct {
		A         []int
		B         []int
		expectedA int
		expectedB int
	}{
		{[]int{5, 6, 7}, []int{3, 6, 10}, 1, 1},
	}

	for _, test := range tests {
		a, b := compareScores(test.A, test.B)
		if a != test.expectedA && b != test.expectedB {
			t.Fatalf("Received: %d, %d Expected: %d, %d", a, b, test.expectedA, test.expectedB)
		}
	}
}
