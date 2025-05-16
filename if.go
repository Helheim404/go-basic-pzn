package main

import "fmt"

func main() {
	name := "Helheim"
	if name == "Helheim" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("who are you?")
	}
	// short statement
	if nama := "gedomazo"; len(nama) < 5 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
