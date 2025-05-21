package main

import "fmt"

// how to declare interface
// now, how to call interface? make a method with same contract with interface. but first make a struct
type HasName interface {
	GetName() string
}

func (person Person) GetName() string {
	return person.Name
}

type Person struct {
	Name string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"Eki"}
	SayHello(person)
	animal := Animal{"Dog"}
	SayHello(animal)
}
