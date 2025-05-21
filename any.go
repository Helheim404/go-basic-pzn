package main

import "fmt"

// empty interface is an any,
func Ups() any {
	// any can be anything, either boolean, int number or even string
	return "Kosong?"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
