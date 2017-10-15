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

	if n < 1 {
		log.Fatal("Input must be greater than 0")
	}

	var values []int
	for i := 0; i <= n-1; i++ {
		var n int
		scan(&n)
		// fmt.Println(n)
		values = append(values, n)
	}

	count := make(map[int]int)
	for _, n := range values {
		count[n]++
	}

	for k, v := range count {
		if v == 1 {
			fmt.Println(k)
		}
	}
}

func scan(n *int) int {
	fmt.Scan(n)
	return *n
}
