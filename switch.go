package main

import (
	"fmt"
)

func main() {
	// switch example
	name := "helheim"

	switch name {
	case "helheim":
		fmt.Println(true, name)

	case "budi":
		fmt.Println(true, name)

	default:
		fmt.Println("who are you?")
	}
	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("panjang nama sudah benar")
	}

	// switch without condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	case length < 5:
		fmt.Println("nama pendek")
	}
}
