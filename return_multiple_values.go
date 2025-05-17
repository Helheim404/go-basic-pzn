package main

import "fmt"

func getFullName() (string, string) {
	return "marshall", "mathers"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)
	// use "_" to ignore value
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
