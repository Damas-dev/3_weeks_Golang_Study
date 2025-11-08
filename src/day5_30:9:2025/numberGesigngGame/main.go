//? Question: Number Guessing Game

//? Write a Go program that lets the user guess a randomly generated number between 0 and 100.

/* Requirements:
*The program should:
~Generate a random number between 0 and 100 using the math/rand package.

~Ask the user to enter their name and then start the game with a friendly greeting.

~Prompt the user to guess the number.

*After each guess:

~If the guess is too high, display: "Too high! Try again."

~If the guess is too low, display: "Too low! Try again."

~If the guess is correct, display: "Congratulations <name>! You guessed it right in X attempts!"

~Use a function named checkGuess(secret, guess int) string that:

~Returns "high", "low", or "correct" depending on the comparison.

~Count and display how many guesses the player made before getting the right number.

*When the user wins, ask:
~"Do you want to play again? (yes/no)"

~If yes → restart the game.

~If no → exit the program with a goodbye message.

*/

/*
	Extra Challenge (optional):

~ Limit the player to 10 attempts.
~ If they fail, display: "Game over! The correct number was X."

~ Add difficulty levels (Easy: 0–50, Hard: 0–100, Expert: 0–500).
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func checkGuess(secret, guess int) string {
	if guess > secret {
		return "high"

	} else if guess < secret {
		return "low"
	} else {
		return "win"
	}

}
func goingOn() string {
	var goingOn string
	fmt.Print("Do you want to play again? (yes/no):")
	fmt.Scan(&goingOn)
	return goingOn
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var name string

	var answer int

	var guessednum int
	fmt.Println(answer)
	fmt.Println("\n\n\n===================== Guess The Number Game  =====================")

	fmt.Print("\n\nWelcome, Please Enter your name to get started,Name:")
	fmt.Scan(&name)

	fmt.Println("\nhello,", name)

	fmt.Println("\n\nKey: [0 ,100) is a range of number form 0 to 99, meaning 0 included while 100 is excluded in the range")

	for i := 1; i <= 11; i++ {
		if i == 11 {
			fmt.Println("Game over! The correct number was", answer)
			if goingOn() == "yes" {
				i = 0
			} else {
				fmt.Println("Good game, goodbye")
				break
			}

		}
		if i == 1 {
			answer = rand.Intn(1000)
			fmt.Print("\nin range [0,100),Guaess the number:")
			fmt.Scan(&guessednum)

		} else {
			fmt.Print("Try again:")
			fmt.Scan(&guessednum)
		}

		status := checkGuess(answer, guessednum)
		if status == "high" {
			fmt.Println("Too high!, ")
			continue
		} else if status == "low" {
			fmt.Println("Too low!, ")
			continue
		} else {
			fmt.Println("Congratulations!", name, " You guessed it right in", i, "attempts!")
			if status := goingOn(); status == "yes" {
				i = 0
			} else {
				fmt.Println("Good game! Goodbye")
				break
			}
		}

	}

}
