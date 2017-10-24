package main

import (
	"fmt"
	"log"
)

func main() {
	var T int
	_, err := fmt.Scan(&T)
	if err != nil {
		log.Fatal(err)
	}

	if T <= 1 {
		log.Fatal("Input must be greater than 1")
	}

	var words []string
	for i := 0; i <= T-1; i++ {
		var w string
		fmt.Scan(&w)
		words = append(words, w)
	}

	for _, word := range words {
		e, o := extractEvenOdd(word)
		fmt.Println(e, o)
	}
}

func extractEvenOdd(word string) (e string, o string) {
	for j, c := range word {
		if j%2 == 0 {
			e = e + string(c)
		} else {
			o = o + string(c)
		}
	}

	return e, o
}
