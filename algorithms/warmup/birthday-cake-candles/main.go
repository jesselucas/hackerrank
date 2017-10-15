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

	if n < 0 {
		log.Fatal("Input must be greater than 0")
	}

	var candles []int
	for i := 0; i <= n-1; i++ {
		var n int
		scan(&n)
		candles = append(candles, n)
	}

	blown := howManyBlown(candles)

	fmt.Println(blown)
}

func howManyBlown(candles []int) int {
	var blown int
	max := candles[0]
	for _, n := range candles {
		if n > max {
			max = n
		}
	}

	// Now count how many max values are in candles slice
	for _, n := range candles {
		if n == max {
			blown++
		}
	}

	return blown
}

func scan(n *int) int {
	fmt.Scan(n)
	return *n
}
