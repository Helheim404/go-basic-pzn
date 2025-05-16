package main

import "fmt"

func main() {
	// ecample of continue. continue simply like skipping some functions with set condition(s)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd loop ", i)
	}
}
