package main

import "fmt"

func main() {
	fmt.Println("This is a summary of Go Arrays")

	//^ About Arrays
	//* Arrays store multiple pieces of the same kind of data.
	//* Data is stored consecutively (in a block of memory).
	//* Each piece of data is called an 'element'.
	//* Arrays are **fixed-size** and cannot be resized after creation.

	//--- Creation & Access ---

	//& Array Creation Types
	// The size is part of the array's type, e.g., '[3]int'

	// 1. Declaration with default values (Zero Values)
	// Elements not set manually will have their default value (e.g., 0 for int).
	var myArray1 [3]int
	fmt.Println("Array 1 (Zero Values):", myArray1) // Output: [0 0 0]

	// 2. Declaration and assignment (Explicit size)
	myArray2 := [3]int{7, 8, 9}
	fmt.Println("Array 2 (Explicit size):", myArray2)

	// 3. Declaration and assignment (Inferred size using '...')
	myArray3 := [...]int{7, 8, 9} // The '...' tells the compiler to count the elements
	fmt.Println("Array 3 (Inferred size):", myArray3)

	// 4. Partial assignment (Elements not listed get default values)
	myArray4 := [4]int{7, 8, 9} // The 4th element will be 0
	fmt.Println("Array 4 (Partial assignment):", myArray4)

	//& Accessing and Modifying Elements
	// Array indexing starts at 0 (the first item is at index 0).

	// Reading an element
	item1 := myArray2[0] // Access using the index.
	fmt.Println("First item (index 0) of Array 2:", item1)

	// Writing/Modifying an element
	myArray1[0] = 7
	myArray1[1] = 8
	myArray1[2] = 9
	fmt.Println("Array 1 (Modified):", myArray1)

	//--- Iteration and Bounds ---

	//& Iteration (Looping through the array)
	// We use the len() function to get the array's size for the loop condition.
	myArray5 := [...]int{10, 20, 30}

	// Using a traditional 'for' loop
	for i := 0; i < len(myArray5); i++ {
		// It's good practice to assign the element to a variable for readability.
		item := myArray5[i]
		fmt.Println("Element at index", i, ":", item)
	}

	// NOTE: You can also use the 'for range' loop (not in the source, but common for arrays/slices):
	// for index, value := range myArray5 { ... }

	//& Bounds Checking (Accessing outside of the array)
	// Attempting to access an element outside the valid index range results in an error.

	//~ Example 1: Compile Time Error
	// Trying to assign a value to myArray[3] when size is 3 (valid indices are 0, 1, 2).
	// var myArray [3]int; myArray[3] = 10 // This causes a compiler error.

	//~ Example 2: Run Time Error
	// Trying to access an element beyond the bounds in a loop.
	// for i := 0; i < 10; i++ { fmt.Println(myArray5[i]) } // This would cause a runtime error when i >= 3.

	//* Recap: Arrays are fixed-size collections. Access is done by index (starting at 0). Out-of-bounds access is an error.
}
