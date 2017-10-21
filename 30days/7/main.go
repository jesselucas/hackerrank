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

	if n <= 1 || n > 1000 {
		log.Fatal("Input must be greater than 1 and less than 1000")
	}

	var a []int
	for i := 0; i <= n-1; i++ {
		var n int
		scan(&n)
		a = append(a, n)
	}

	fmt.Println(sprintReverse(a))
}

func sprintReverse(a []int) string {
	out := ""
	for i := len(a) - 1; i >= 0; i-- {
		n := a[i]
		out += fmt.Sprint(n)

		// Add space to all elements except last
		if i != 0 {
			out += " "
		}
	}
	return (out)
}

func scan(n *int) int {
	fmt.Scan(n)
	return *n
}
