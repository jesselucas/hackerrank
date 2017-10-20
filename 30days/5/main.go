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

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, (n * i))
	}
}
