package main

import "fmt"

func main() {
	counter := 0

	// in golang, function can access variable outside the func
	increment := func() {
		fmt.Println("increment")
		counter++
	}
	// increment() is a closure, this feature is useful but do it with moderation beccuse it can be confusing when your block code is already so long
	increment()
	increment()
	fmt.Println(counter)
}
