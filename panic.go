package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		// panic is for stopping the program completely
		panic("Whoops, something went wrong")
	}
}

func main() {
	runApp(true)
}
