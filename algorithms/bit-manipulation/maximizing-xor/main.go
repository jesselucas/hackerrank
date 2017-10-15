package main

import (
	"fmt"
	"log"
)

func main() {
	var l, r int
	_, err := fmt.Scan(&l, &r)
	if err != nil {
		log.Fatal(err)
	}

	if l > r {
		log.Fatalf("left: %d is greater than right: %d", l, r)
	}

	var max int
	for i := l; i <= r; i++ {
		for j := i; j <= r; j++ {
			xor := i ^ j
			if xor > max {
				max = xor
			}
		}
	}

	fmt.Println(max)
}
