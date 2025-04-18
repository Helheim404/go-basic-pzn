package main

import "fmt"

func main() {
	// old way of declaring array
	var names [3]string

	names[0] = "Helheim"
	names[1] = "404"
	names[2] = "Goddess of Helheim"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// another way of declaring array
	// var values = [3]int{
	var values = [...]int{
		1,
		2,
		3,
	}
	fmt.Println(values)

	// array function
	fmt.Println(len(values))

}
