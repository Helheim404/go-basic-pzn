package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("you cannot divide by zero")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err != nil {
		fmt.Println("Errors", err.Error())
	} else {
		fmt.Println("hasil", hasil)
	}
}
