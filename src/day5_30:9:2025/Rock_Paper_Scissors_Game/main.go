//? Question: Rock, Paper, Scissors Game

//? Write a Go program that allows the user to play Rock, Paper, Scissors against the computer.

/* Requirements:
* The program should:
~ Ask the user to enter their name.
~ Greet the user and explain the rules (rock beats scissors, scissors beats paper, paper beats rock).
~ Prompt the user to choose between "rock", "paper", or "scissors".
~ The computer should randomly choose one of the three options.
~ Compare both choices and determine the winner:
   - Rock beats Scissors
   - Scissors beats Paper
   - Paper beats Rock
   - If both are the same, it’s a draw.
~ Display the result of each round (who won).
~ Keep track of the score (user wins, computer wins, draws).
~ Ask the user if they want to play again (yes/no).
~ If "yes" → play another round.
~ If "no" → display the final score and end the game with a goodbye message.

*/

/*
	Extra Challenge (optional):

~ Let the user choose how many rounds they want to play before starting.
~ Display the winner of the whole game at the end (best of N rounds).
~ Handle invalid inputs gracefully (e.g., if the user types “rok” instead of “rock”).
~ Add difficulty levels (Easy: computer random, Hard: computer tries to win more).
*/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// checkWinner compares user and computer choices and returns the result
func checkWinner(userChoice, compChoice string) string {
	if userChoice == compChoice {
		return "draw"
	}

	if (userChoice == "rock" && compChoice == "scissors") ||
		(userChoice == "scissors" && compChoice == "paper") ||
		(userChoice == "paper" && compChoice == "rock") {
		return "user"
	}

	return "computer"
}

// playAgain asks the user if they want to play again
func playAgain() bool {
	var response string
	fmt.Print("Do you want to play again? (yes/no): ")
	fmt.Scan(&response)
	response = strings.ToLower(response)
	return response == "yes"
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var name string
	fmt.Println("===================== Rock, Paper, Scissors Game =====================")
	fmt.Print("\nWelcome! Please enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("\nHello, %s! Here are the rules:\n", name)
	fmt.Println("- Rock beats Scissors")
	fmt.Println("- Scissors beats Paper")
	fmt.Println("- Paper beats Rock")

	choices := []string{"rock", "paper", "scissors"}
	var userChoice string
	userWins, compWins, draws := 0, 0, 0

	for {
		fmt.Print("Choose (rock/paper/scissors): ")
		fmt.Scan(&userChoice)
		userChoice = strings.ToLower(userChoice)

		// Validate user input
		if userChoice != "rock" && userChoice != "paper" && userChoice != "scissors" {
			fmt.Println("Invalid choice! Please type rock, paper, or scissors.")
			continue
		}

		// Computer choice
		compChoice := choices[rand.Intn(len(choices))]
		fmt.Println("Computer chose:", compChoice)

		// result
		result := checkWinner(userChoice, compChoice)

		switch result {
		case "user":
			fmt.Println("You win this round!")
			userWins++
		case "computer":
			fmt.Println("Computer wins this round!")
			compWins++
		default:
			fmt.Println("It's a draw!")
			draws++
		}

		fmt.Printf("\nScoreboard: %s=%d | Computer=%d | Draws=%d\n\n", name, userWins, compWins, draws)

		if !playAgain() {
			fmt.Println("\n===================== Final Results =====================")
			fmt.Printf("%s: %d wins | Computer: %d wins | Draws: %d\n", name, userWins, compWins, draws)

			if userWins > compWins {
				fmt.Printf("Congratulations %s, you are the overall winner!\n", name)
			} else if compWins > userWins {
				fmt.Println("The computer wins this time. Better luck next round!")
			} else {
				fmt.Println("It's a tie overall!")
			}
			fmt.Println("Goodbye!")
			break
		}
	}
}
