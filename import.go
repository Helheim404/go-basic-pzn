package main

import (
	"fmt"
	"golang-hello-world/helper"
)

func main() {
	result := helper.SayHello("warudo")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // cannot be accessed because version is using lowecase. To be able access it, you must declared with uppercase in the first alphabetL
}
