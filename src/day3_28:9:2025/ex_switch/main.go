package main

import "fmt"

//?--Summary:
//?  Create a program to display a classification based on age.
//?
//?--Requirements:
//* Use a `switch` statement to print the following:
//?  - "newborn" when age is 0
//?  - "toddler" when age is 1, 2, or 3
//?  - "child" when age is 4 through 12
//?  - "teenager" when age is 13 through 17
//?  - "adult" when age is 18+

func main() {
	var age uint
	fmt.Print("Please enter your age(in years): ")
	fmt.Scan(&age)

	switch {
	//"newborn" when age is 0
	case age == 0:
		fmt.Println("newborn")

	//"toddler" when age is 1, 2, or 3
	case age <= 3:
		fmt.Println("Toddler")

	//"child" when age is 4 through 12
	case age <= 12:
		fmt.Println("Child")

	//"teenager" when age is 13 through 17
	case age <= 17:
		fmt.Println("Teenager")

	//"adult" when age is 18+
	case age >= 18:
		fmt.Println("Adult")

	default:
		fmt.Println("Error")

	}

}
