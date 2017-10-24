package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var inputI int
	var inputD float64
	var inputS string

	// Read and save an integer, double, and String to your variables.
	for scanner.Scan() {
		var err error

		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			inputI = i
			continue
		}

		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err == nil {
			inputD = f
			continue
		}

		s := scanner.Text()
		if s != "" {
			inputS = s
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Print the sum of both integer variables on a new line.
	fmt.Println(int(i) + inputI)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+inputD)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Printf("%s%s\n", s, inputS)
}
