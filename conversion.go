package main

import "fmt"

func main() {
	// this is how you convert number data type
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	// this function gonna trigger number overflow, max value of int16 is 32767, so when you convert 32768 to int16, it's gonna overflow to the lowest value which is -32768
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// this is how you convert string data type
	var name = "Helheim"
	// value of e is index 0 of variable name which is "H" or in interger is 72
	var e uint8 = name[0]
	// value of eString variable is "H" which is a string conversion of variable e with value 72
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
