package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println("Hello, World.")
	fmt.Println(input)
}
