package main

import "testing"

func Test_minMaxXOR(t *testing.T) {
	tests := []struct {
		nums     []uint64
		expected uint64
	}{
		{[]uint64{1, 2, 3}, 2},
		{[]uint64{1001, 1002, 1003, 1004}, 5},
		{[]uint64{2, 2, 3}, 1},
		{[]uint64{2, 2, 2}, 0},
		{[]uint64{0, 1, 3, 4, 6}, 4},
		{[]uint64{12, 8, 9, 11, 14}, 4},
	}

	for i, test := range tests {
		xor := minMaxXOR(test.nums)
		if xor != test.expected {
			t.Fatalf("Test Failed: %d, Received: %d, Expected: %d", i, xor, test.expected)
		}
	}
}
