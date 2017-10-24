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
	fmt.Println(amIWeird(n))
}

func amIWeird(n int) string {
	if n%2 == 1 {
		return "Weird"
	}

	if n >= 6 && n <= 20 {
		return "Weird"
	}

	return "Not Weird"
}
