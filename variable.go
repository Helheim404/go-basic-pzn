package main

import "fmt"

func main() {
	var name string //string type variable

	name = "Helheim 404"
	fmt.Println(name)

	name = "Goddess of Helheim"
	fmt.Println(name) // this should be changed from "Helheim 404" to "Goddess of Helheim" because code is read from top to bottom
}
