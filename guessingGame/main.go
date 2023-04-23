package main

import (
	"fmt"
	"math/rand"
	"os"
)

func game(file *os.File) {
	fmt.Println("Welcome to the guessing game! Guess the number between 1 and 100.")

	target := rand.Intn(100) + 1

	fmt.Println(target)

	var guess int
	var guesses int

	var name string

	fmt.Println("Please enter your name: ")
	fmt.Scan(&name)

	for guess != target {
		fmt.Println("Please enter your guess: ")
		fmt.Scan(&guess)
		guesses++
		if guess > target {
			fmt.Println("Too high!")
		} else if guess < target {
			fmt.Println("Too low!")
		} else {
			break
		}
	}
	fmt.Println("You got it! It took you", guesses, "guesses.")
	fmt.Fprintf(file, "Player name: %s, Score: %d, Number: %d\n", name, guesses, target)
}

func main() {
	file, err := os.OpenFile("scores.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	for {
		game(file)

		var playAgain string
		fmt.Println("Would you like to play again? (y/n)")
		fmt.Scan(&playAgain)

		if playAgain != "y" {
			fmt.Println("Thanks for playing! :) ")
			break
		}
	}
}
