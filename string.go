package main

import "fmt"

func main() {
	fmt.Println("this is how you use string in golang using double quotes")
	//function in string, which len("string") and "string"[number]. len for length and number for byte
	//code below the length which is 56 include space
	fmt.Println(len("this is how you use string in golang using double quotes"))
	//code below going to print "t" as a byte and "h" as a byte
	fmt.Println("this is how you use string in golang using double quotes"[0])
	fmt.Println("this is how you use string in golang using double quotes"[1])
}
