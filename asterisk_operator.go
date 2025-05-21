package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Cilacap", "Central Java", "Indonesia"}
	address2 := &address1
	address2.City = "Purwokerto"
	fmt.Println(address1)
	fmt.Println(address2)

	//	address2 = &Address{"Surabaya", "East Java", "Indonesia"}

	*address2 = Address{"Surabaya", "East Java", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
