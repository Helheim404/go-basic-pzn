package main

import (
	"fmt"
	"golang-hello-world/database"
	_ "golang-hello-world/internal" // _ used for package that imported to run without being used in main function for init function
)

func main() {
	fmt.Println(database.GetDatabase())
}
