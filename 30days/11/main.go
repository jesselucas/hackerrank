package main

import "fmt"
import "math"

func main() {
	n := 6
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

	fmt.Println(hourglassMaxSum(matrix))
}

func hourglassMaxSum(matrix [][]int) int {
	hourglass := [][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 1},
		[]int{2, 0},
		[]int{2, 1},
		[]int{2, 2},
	}

	maxSum := math.MinInt16
	for i := 0; i <= len(matrix)-3; i++ {
		r := matrix[i]
		for j := 0; j <= len(r)-3; j++ {
			// reset sum
			sum := 0

			// Check hour glass
			for _, p := range hourglass {
				sum += matrix[p[0]+i][p[1]+j]
			}

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}
