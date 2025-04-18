package main

import (
	"fmt"
)

func main() {

	// slice is simply a dynamic array, can be referenced from whole array or part of array
	// slice can be declare by array[lower index : higher index] example array[0:5], or array[:4] for slice from index 0 to index before 4, or array[0:] for slice from index 0 to last index or array[:] for slice from whole array
	names := [...]string{"Hel", "Goddess", "of", "Helheim"}

	// slice will be referenced from index 1 to index before 4 of array names
	slice1 := names[1:4]

	fmt.Println(slice1)

	// function in slice
	// len(slice) to determine the length of slice
	// cap(slice) to check the capacity of array from first index of slice
	// append(slice, value) to add value to slice and if the capacity of array is not enough, it will make new array and new capacity, BUT it's not accessible outside of new slice
	// make([]TypeData, length, capacity) to make new slice with length and capacity
	// copy(slice1, slice2) to copy slice2 to slice1

	// append slice
	// array of days
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	// slice of days from index 4 to last
	fmt.Println(days)
	daySlice1 := days[4:]
	// changing value of index 0 and 1 in array days from dayslice1
	daySlice1[0] = "Jumat Baru"
	daySlice1[1] = "Sabtu Baru"
	daySlice1[2] = "Minggu Baru"
	// index 4 and 5 of days is now should be Jumat Baru and Sabtu Baru
	fmt.Println(days)
	// now let's try append slice
	daySlice2 := append(daySlice1, "Senin Baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// make slice
	newSlice := make([]string, 2, 5) // why the length and capacity is different? because when we add value to slice, there is still capacity for new value so the program won't be bloated with new array when we added more value in the future
	newSlice[0] = "gedomazo"
	newSlice[1] = "gedomazo"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// to add more value to slice use append
	newSlice2 := append(newSlice, "gedomazo")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// difference from array and slice
	thisArray := [...]int{1, 2, 3}
	thisSlice := []int{1, 2, 3}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
