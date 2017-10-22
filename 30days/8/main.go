package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read stdin
	var i int
	var n int
	var err error

	phoneBook := make(map[string]string)
	for scanner.Scan() {
		if i == 0 {
			n, err = strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
		}

		if i > 0 && i <= n {
			// Break string into map
			p := strings.Split(scanner.Text(), " ")
			phoneBook[p[0]] = p[1]
		} else if i > n {
			v, ok := phoneBook[scanner.Text()]
			if ok {
				fmt.Printf("%s=%s\n", scanner.Text(), v)
			} else {
				fmt.Println("Not found")
			}
		}

		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
