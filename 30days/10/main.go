package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(consecutiveOnes(n))
}

func consecutiveOnes(n int) int {
	max := 0
	cOnes := 0
	previousBit := 0
	for n > 0 {
		r := n % 2

		// Make sure we count the first 1
		if r == 1 && cOnes == 0 {
			cOnes++
		}

		// Count consecutive 1s
		if previousBit == 1 {
			cOnes++
		}
		// Reset to 0 if we find a 0
		if r == 0 {
			cOnes = 0
		}

		// Store previous bit
		previousBit = r

		// Check max
		if cOnes > max {
			max = cOnes
		}

		// Reduce n
		n = n / 2
	}

	return max
}
