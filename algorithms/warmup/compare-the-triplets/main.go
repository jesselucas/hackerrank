package main

import (
	"fmt"
	"log"
)

func main() {
	var A []int
	var B []int

	// Scan Alice then Bog
	for i := 0; i <= 5; i++ {
		var n int
		if err := scan(&n); err != nil {
			log.Fatal(err)
		}
		if i < 3 { // scan Alice
			A = append(A, n)
		} else { // scan Bob
			B = append(B, n)
		}
	}

	fmt.Println(compareScores(A, B))
}

func compareScores(A, B []int) (int, int) {
	scoreA := 0
	scoreB := 0
	for i, a := range A {
		b := B[i]

		if a > b {
			scoreA++
		} else if a < b {
			scoreB++
		}
	}

	return scoreA, scoreB
}

func scan(n *int) error {
	_, err := fmt.Scan(n)
	if err != nil {
		return err
	}

	return nil
}
