package main

import "fmt"

func main() {
	// two way of declaring variable
	var name string

	name = "Helheim 404"
	fmt.Println(name)

	name = "Goddess of Helheim"
	fmt.Println(name) // this should be changed from "Helheim 404" to "Goddess of Helheim" because code is read from top to bottom

	// another way of declaring variable
	var username = "Helheim"
	fmt.Println(username)

	// var actually doesn't alwaysneed to be declared, but it's a good practice to declare it
	// you can initialize the variable without var when you also declare it using " := " but only at the first time you declare it.
	// like this code below
	githubUsername := "Helheim"
	fmt.Println(githubUsername)

	githubUsername = "Goddess of Helheim"
	fmt.Println(githubUsername)

	// golang can make you declare multiple variables at once (nice) like this example below
	var (
		firstName = "Hello"
		lastName  = "World"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
