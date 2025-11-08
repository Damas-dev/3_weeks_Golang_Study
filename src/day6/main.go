package main

import "fmt"

func main() {
	fmt.Println("This is a summary of Go Structs")

	//^ What are Structures (Structs)?
	//* Structs let you store related data together in groups.
	//* They are similar to a "class" in other programming languages.
	//* Each piece of data inside a struct is called a 'field'.
	//* They help organize code and data efficiently.

	//--- Defining a Struct ---
	// We use the 'type' keyword to define a structure.
	type Sample struct {
		field string // A field of type string
		a, b  int    // Two fields of type int
	}

	//--- Creating (Instantiating) a Struct ---

	//& 1. Instantiation (Creating an instance of the struct)
	data1 := Sample{"word", 1, 2} // Shorthand, order matters
	fmt.Println("Data 1 (Shorthand):", data1)

	//& 2. Instantiation (Using Field Names - Best Practice)
	data2 := Sample{
		field: "word", // Explicitly naming fields
		a:     1,
		b:     2,
	}
	fmt.Println("Data 2 (Named fields):", data2)

	//& 3. Default Values (Zero Values)
	// If you don't set a field, it gets its default (zero) value.
	// string gets "", int gets 0.
	data3 := Sample{a: 5} // 'field' will be "" and 'b' will be 0
	fmt.Println("Data 3 (Default values):", data3)

	//--- Accessing and Changing Fields ---

	//& Accessing Fields (Reading)
	word := data2.field // Access using dot-notation
	val_a := data2.a    // Reading the 'a' field
	fmt.Println("Accessed field:", word, "and a:", val_a)

	//& Modifying Fields (Writing)
	data2.field = "hello" // Fields can be written to
	data2.a = 10          // Changing the value of 'a'
	data2.b = 20          // Changing the value of 'b'
	fmt.Println("Data 2 (After modification):", data2)

	//--- Anonymous Structures ---

	//& Anonymous/Inline Structs
	// These are created right where you need them, often inside a function.
	// Useful for temporary data or working with library functions.

	// 1. Anonymous Struct with Default Values (using 'var')
	var anonymousSample1 struct { // No name (anonymous)
		field string // Defining the fields inline
		a, b  int
	}
	anonymousSample1.field = "hello" // Fields have default values until set
	anonymousSample1.a = 9
	fmt.Println("Anonymous Struct 1:", anonymousSample1)

	// 2. Anonymous Struct (Shorthand and Initialization)
	anonymousSample2 := struct {
		field string
		a, b  int
	}{
		"hello", // Must define a value for *every* field in this shorthand
		1, 2,
	}
	fmt.Println("Anonymous Struct 2:", anonymousSample2)

	//* Recap: Structs group data points called fields.
	//* Access is done with dot-notation.
}
