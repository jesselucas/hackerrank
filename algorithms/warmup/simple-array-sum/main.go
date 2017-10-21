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

	var a []int
	for i := 0; i <= n-1; i++ {
		var n int
		scan(&n)
		a = append(a, n)
	}

	fmt.Println(sumInts(a))
}

func sumInts(a []int) int {
	sum := 0
	for _, n := range a {
		sum += n
	}

	return sum
}

func scan(n *int) int {
	fmt.Scan(n)
	return *n
}
