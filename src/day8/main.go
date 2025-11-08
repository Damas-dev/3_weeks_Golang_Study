package main

import "fmt"

func main() {
	fmt.Println("This is a summary of Go Slices")

	//^ What are Slices?
	//* Slices are companion types that work with arrays.
	//* They create a "view" into an array.
	//* They are **dynamic** (resizable) and NOT fixed in size.
	//* Slices always require an underlying array.

	//--- Creation & Access ---

	//& 1. Short Declaration (creates a slice and an underlying array)
	mySlice := []int{1, 2, 3}
	fmt.Println("Slice 1:", mySlice)

	//& 2. Accessing Elements
	// Access is the same as an array (using index).
	item1 := mySlice[0]
	fmt.Println("First item:", item1)

	//& 3. Preallocation (using make())
	// Preallocates a slice with a capacity, useful when size is known but values aren't.
	preallocated := make([]int, 10)
	fmt.Println("Preallocated Slice (10 zeros):", preallocated)

	//--- Slicing Syntax ---

	//& Slicing (Creating a new slice from an array or existing slice)
	numbers := [...]int{1, 2, 3, 4, 5} // An array to start with

	// Syntax is [start (inclusive) : end (exclusive)]
	sliceFromStart := numbers[:2] // Start at 0, up to (but not including) index 2 -> [1, 2]
	sliceToEnd := numbers[3:]     // Start at index 3, go to the end -> [4, 5]

	fmt.Println("Sliced (0:2):", sliceFromStart)
	fmt.Println("Sliced (3:end):", sliceToEnd)

	//--- Resizing ---

	//& Dynamic Resizing (using append())
	// The append() function adds elements and resizes the slice.
	mySlice = append(mySlice, 4, 5, 6)
	fmt.Println("Slice after append:", mySlice) // [1, 2, 3, 4, 5, 6]

	//& Appending one slice to another
	part1 := []int{1, 2}
	part2 := []int{3, 4}
	// Use the "three dots" (...) to unpack the second slice.
	combined := append(part1, part2...)
	fmt.Println("Combined Slices:", combined)

	//* Recap: Slices are dynamic array views, created with '[]type' or 'make()', and resized with 'append()'.
}
