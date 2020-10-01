// This program randomly takes in a non-negative integer and generates maxTurns amount of turns
// And tries to get the same number as the user input.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Max amount of times to loop and generate a new random number.
const maxTurns = 5

func main() {
	// Generate a random number by creating unique seeds
	rand.Seed(time.Now().UnixNano())

	// Select the number entered by user
	args := os.Args[1:]

	// If user enters too many numbers display the following message and end program.
	if len(args) != 1 {
		fmt.Println("Too many numbers! Please enter only one number.")
		return
	}

	// If user enters a non number character display the following message and end program.
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number!")
		return
	}

	// If user enters a negative number display the following message and end program.
	if guess < 0 {
		fmt.Println("Please enter a positive integer")
		return
	}

	// For clarity print out the number the user inputted.
	fmt.Println("Your number is ", guess)

	// Produce a random number
	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		fmt.Printf("The random number is: %d\n", n)

		// User wins when guess == random number.
		if n == guess {
			fmt.Println("You win!!!")
			return
		}
	}
	// Otherwise if the numbers didn't match, print out this message.
	fmt.Println("Better luck next time...")
}
