package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("eminem")
	fmt.Println(result)
	// or
	fmt.Println(getHello("master"))
	fmt.Println(getHello("marshall"))
}
