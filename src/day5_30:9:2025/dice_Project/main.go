/* --Summary:
~ Create a program that can perform dice rolls. The program should
~ report the results as detailed in the requirements.
*/
/* --Requirements:
* Print the sum of the dice roll
* Print additional information in these cirsumstances:
~"Snake eyes": when the total roll is 2, and total dice is 2
~"Lucky 7": when the total roll is 7
~"Even": when the total roll is even
~"Odd": when the total roll is odd
*/
//* The program must handle any number of dice, rolls, and sides

//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var dice, sides, rolls int
	fmt.Println("\n--- Dice Rolling Game ---")
	// user for input
	fmt.Print("Enter number of dice: ")
	fmt.Scan(&dice)

	fmt.Print("Enter number of sides per dice: ")
	fmt.Scan(&sides)

	fmt.Print("Enter number of rolls: ")
	fmt.Scan(&rolls)

	for i := 1; i <= rolls; i++ {
		total := 0
		fmt.Printf("\nRoll %d:\n", i)

		for d := 1; d <= dice; d++ {
			roll := rand.Intn(sides) + 1
			fmt.Printf("  Dice %d: %d\n", d, roll)
			total += roll
		}

		fmt.Printf("Total Roll: %d\n", total)

		if dice == 2 && total == 2 {
			fmt.Println("Snake eyes!")
		}
		if total == 7 {
			fmt.Println("Lucky 7!")
		}
		if total%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
