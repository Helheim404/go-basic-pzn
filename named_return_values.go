package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "marshall"
	middleName = "mathers"
	lastName = "III"

	return firstName, middleName, lastName
}

func main() {
	x, y, z := getCompleteName()
	fmt.Println(x, y, z)
}
