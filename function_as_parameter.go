package main

import "fmt"

// filter is a func inside func as a parameter. it can be  used as long as there is anoter function with the same contract
// this filter funct is too long, in Go we can make Type declaration or alias for this filter func
type Filter func(string) string

// long and tedious way to declarae func as parameter (filter func)
// func sayHelloWithFilter(name string, filter func(string) string)
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

// this function is the same contract fith filter parameter, which is function with string parameter and returning string value
func spamFilter(name string) string {
	if name == "Anjing" {
		return "#censored#"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Marshall", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
