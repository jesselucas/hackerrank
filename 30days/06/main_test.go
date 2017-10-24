package main

import "testing"

func Test_extractEvenOdd(t *testing.T) {
	tests := []struct {
		word string
		even string
		odd  string
	}{
		{"Hacker", "Hce", "akr"},
		{"Rank", "Rn", "ak"},
	}

	for _, test := range tests {
		e, o := extractEvenOdd(test.word)
		if e != test.even && o != test.odd {
			t.Fatal()
		}
	}
}
