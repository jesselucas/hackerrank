package main

import "testing"

func Test_amIWeird(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{3, "Weird"},
		{2, "Not Weird"},
		{4, "Not Weird"},
		{6, "Weird"},
		{8, "Weird"},
		{10, "Weird"},
		{14, "Weird"},
		{20, "Weird"},
		{24, "Not Weird"},
		{100, "Not Weird"},
	}

	for _, test := range tests {
		out := amIWeird(test.n)
		if out != test.expected {
			t.Fatal("Received: ", out, "Expected: ", test.expected)
		}
	}
}
