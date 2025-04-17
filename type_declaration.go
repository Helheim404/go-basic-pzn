package main

import "fmt"

func main() {
	// type declaration is a way to declare a variable with a specific data type. In this case string is declared as NoKTP
	type NoKTP string
	// var ktpku (variable name) NoKTP (data type which is string) = "123456789" (value)
	var ktpKu NoKTP = "123456789"
	var contohKtp = NoKTP(ktpKu)
	// output is string
	fmt.Println(ktpKu)
	fmt.Println(NoKTP("22222222"))
	fmt.Println(contohKtp)

}
