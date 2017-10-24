package main

import (
	"fmt"
	"log"
	"math"
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

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = []int{}
	}

	j := 0
	for i := 0; i < (n * n); i++ {
		var v int
		fmt.Scan(&v)
		if matrix[j] == nil {
			matrix[j] = []int{}
		}
		matrix[j] = append(matrix[j], v)
		if (i+1)%n == 0 {
			j++
		}
	}

	fmt.Println(absDiagonalDif(matrix))
}

func absDiagonalDif(matrix [][]int) float64 {
	// Diagonal
	length := len(matrix)
	da := 0
	db := 0
	for i := 0; i < len(matrix); i++ {
		da += matrix[i][i]
		db += matrix[i][length-1-i]
	}
	a := da - db
	return math.Abs(float64(a))
}
