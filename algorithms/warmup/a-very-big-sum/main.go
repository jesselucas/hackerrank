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
		log.Fatal("Input must be greater than 1")
	}

	var sum int
	for i := 0; i <= n-1; i++ {
		var n int
		fmt.Scan(&n)
		sum += n
	}

	fmt.Println(sum)
}
