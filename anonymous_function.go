package main

import "fmt"

// for anonymous func, first declare type of func, then make variable as a function holder
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're blocked ", name)
	} else {
		fmt.Println("Welcome, ", name)
	}
}

func main() {
	// this is anonymous func inside variable blacklist
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	// simply called the registerUser func and blacklist variable to execute blacklisted name
	registerUser("Helheim404", blacklist)
	// or called the registerUser func and directly declare the func
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
