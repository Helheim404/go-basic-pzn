package main

import "fmt"

// the three dots is esentially a slice inside a function as a value. it's called variable argument
func sumAll(numbers ...int) int {
	total := 0
	// "_" used for ignoring index
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	// you can acctually insert slice into variable arrgument using this method

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
