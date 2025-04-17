package main

import "fmt"

func main() {
	// constant is a variable that cannot be changed and cannot be redeclared, but it's not error when it's not used
	const firstName string = "Helheim"
	const lastName = "404"

	// const also can be declared for multiple variables at once
	const (
		username = "Goddess of Helheim"
		password = "404"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(username)
	fmt.Println(password)

	// firstName = "Goddess of Helheim"
	// lastName = "Hel of Helheim"

}
