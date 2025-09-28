//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

// ? Write a function that accepts a person's name as a function parameter and displays a greeting to that person.
func greeting(name string) {
	fmt.Println("hello", name, "!")
}

// ? qn2 Write a function that returns any message, and call it from within fmt.Println()
func message() string {
	return "I will become a go master one day"
}

// ? Write a function to add 3 numbers together, supplied as arguments, and return the answer
func add(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// ? Write a function that returns any number
func anyNumber() int {
	num := rand.Intn(101)
	return num
}

// ? Write a function that returns any two numbers
func any2Numbers() (int, int) {
	num1 := rand.Intn(101)
	num2 := rand.Intn(101)

	return num1, num2
}

func main() {
	greeting("Damas") //! qn1

	fmt.Println(message()) //! qn2

	result := add(34, 53, 51)                      //! qn3
	fmt.Println("the sum of 3 numbers is", result) //! qn3

	//? Add three numbers together using any combination of the existing functions. Print the result

	randnum1, randnum2 := any2Numbers()

	samOf3randomNums := add(anyNumber(), randnum1, randnum2)

	fmt.Println("the sum of three random numbers is", samOf3randomNums)
}
