package main

import "testing"

func Test_NewPerson(t *testing.T) {
	tests := []struct {
		initialAge int
		expected   int
	}{
		{-1, 0},
		{10, 10},
	}

	for _, test := range tests {
		p := person{}
		p = p.NewPerson(test.initialAge)
		if p.age != test.expected {
			t.Fatal("Received: ", p.age, "Expected: ", test.expected)
		}
	}
}
