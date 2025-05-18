package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("End Application")
	// correct way to use recover, insert in defer function
	message := recover()
	fmt.Println("Panic is triggered, ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		// panic is for stopping the program completely
		// to catch the error message, we need to use recover
		panic("Whoops, something went wrong")
	}
}

func main() {
	runApp(true)
}
