package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()
	// resultString := result.(string) // type assertion from any into string
	// fmt.Println(resultString)
	// beware when using type assertion, if the type assertion is not right, it will cause panic
	// you can using switch
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("any", value)
	}
}
