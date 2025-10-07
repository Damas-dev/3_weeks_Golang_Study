package main

import "fmt"

func main() {
	fmt.Println("this is Day4")
	//^ loop
	//* for loop (for)

	//example
	for i := 0; i <= 20; i++ { //! i++ is the post statement: done last
		if i%2 == 0 {
			continue //! this key word is use to skip all content under this condition
		}
		fmt.Println(i)
	}

	//* while loop(for)
	//! in while loop we use the word for insted of while

	age := 0

	for age >= 0 { //! this is an infinte loop
		if age == 19 {
			break //! this is used to make the loo finete
		}
		fmt.Println(age)
		age++

	}
	fmt.Println("the above is 18-")

}
