package main

import "fmt"

func main() {
	// for loop
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("loop ", counter)
	// 	counter++
	// }
	fmt.Println("done!!")
	// statement can be omitted in for loops, which init statement and post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("loop ", counter)
	}
	fmt.Println("done!!")

	// for can be used for array, slice and map iteration
	// slice manually accessed
	names := []string{"budi", "ini", "abra"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	// slice accessed using range
	cities := []string{"jakarta", "bandung", "semarang"}
	for index, value := range cities {
		fmt.Println("index", index, "=", value)
	}
	// because in go every variable must be used when accessing data collection and index variable sometimes not needed, then can be replaced with "_"
	for _, value := range cities {
		fmt.Println(value)
	}
}
