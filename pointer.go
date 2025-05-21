package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// pointer in golang is passing by value (by default), so if we send var1 into var2 the value is duplicated, not referenced
func main() {
	// address1 := Address{"Cilacap", "Central Java", "Indonesia"}
	// address2 := address1
	// address2.City = "Purwokerto"
	// // this is example from pointer by value. address2.City is different from address1.City.
	// fmt.Println(address1)
	// fmt.Println(address2)

	// to use pointer by reference, we need to use "&" symbol
	address1 := Address{"Cilacap", "Central Java", "Indonesia"}
	address2 := &address1
	address2.City = "Purwokerto"
	fmt.Println(address1)
	fmt.Println(address2)
}
