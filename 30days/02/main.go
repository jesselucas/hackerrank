package main

import (
	"fmt"
	"log"
)

func main() {
	var meal float64
	var tip, tax float64
	_, err := fmt.Scan(&meal, &tip, &tax)
	if err != nil {
		log.Fatal(err)
	}
	cost := mealCost(meal, tip, tax)
	fmt.Printf("The total meal cost is %d dollars.\n", round(cost))
}

func mealCost(meal, tip, tax float64) float64 {
	return meal + (meal * (tip / 100)) + (meal * (tax / 100))
}

func round(a float64) int {
	if a < 0 {
		return int(a - 0.5)
	}
	return int(a + 0.5)
}
