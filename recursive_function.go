package main

import "fmt"

// recursive function is runction that loop on itlsef, typically used to simplify or to avoid for loop and used in Factorial problem for example.
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// now using recursive functition
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
