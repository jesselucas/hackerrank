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

	fmt.Println(factorial(n))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}
