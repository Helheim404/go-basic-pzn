package main

import "fmt"

func main() {
	// map is a collection of key-value pairs which not like array or slice using number as an index
	// in map index is called key and must be unique. if the key is named with the same name, it will overwrite the previous value
	// data in map can be stored infinitely as long as the key is named differently with each other.

	// map can be declared like this
	var person map[string]string = map[string]string{}
	person["name"] = "Helheim" // this is example of map[key] = value
	person["address"] = "Jakarta"

	fmt.Println(person)
	fmt.Println(person["name"]) // this is example of map[key]
	fmt.Println(person["address"])

	// or map can be declare like this
	person1 := map[string]string{
		"name":    "Gedomazo",
		"address": "Konoha",
	}
	fmt.Println(person1)
	fmt.Println(person1["name"])
	fmt.Println(person1["address"])

	// function that can be used in Map is len(map), map[key], map[key] = value, make(map[TypeKey]TypeValue), delete(map, key)
	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Helheim"
	book["editor"] = "Google"

	delete(book, "wrong")

	fmt.Println(book)
}
