package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Selesai memanggil function")
}

// defer is used to execute code after function is finished regardless of if it's success or error
func runApplication() {
	defer logging()
	fmt.Println("Executing...")
}

func main() {
	runApplication()
}
