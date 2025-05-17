package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// varibale goodbye now have a value of a function getGoodBye. you can call getGoodBye function as goodbye variable using goodbye()
	goodbye := getGoodBye

	fmt.Println(goodbye("briar"))
}
