package main

import "fmt"

// nil can only be used for fuction, interface, map, slice, pointer and channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("HEAD")

	if data == nil {
		fmt.Println("Data is nil")
	} else {
		fmt.Println(data["name"])
	}
}
