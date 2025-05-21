package main

import "fmt"

// this is declaration of struct
type Customer struct {
	Name, Address string
	Age           int
}

// struct also can be called as method.
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	/*struct cannot be called directly, you must initialize it first.*/
	// thi is example how to initialzie struct
	var kidd Customer
	kidd.Name = "Kidd"
	kidd.Address = "Jakarta"
	kidd.Age = 30

	fmt.Println(kidd)
	fmt.Println("Name = ", kidd.Name)
	fmt.Println("Address = ", kidd.Address)
	fmt.Println("Age = ", kidd.Age)

	// another example how to initialize struct

	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     29,
	}

	fmt.Println(joko)

	// or
	budi := Customer{"Budi", "Indonesia", 28}
	fmt.Println(budi)
	// sayHello cannot be called directly, because struct is not a function, but it can be called as a method
	budi.sayHello(joko.Name)
}
