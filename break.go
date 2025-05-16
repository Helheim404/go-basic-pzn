package main

import "fmt"

func main() {
	// example of break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("loop ", i)
	}
}
