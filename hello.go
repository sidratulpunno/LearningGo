package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	randomNumber := rand.Intn(100) + 1

	fmt.Println("I'm thinking of a number between 1 and 100.")

	for {
		fmt.Print("Enter your guess: ")

		// Read a line of input from the user
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		// Remove the newline character
		input = strings.TrimSpace(input)

		// Convert the input to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if guess < randomNumber {
			fmt.Println("Too low!")
		} else if guess > randomNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You guessed it!")
			break
		}
	}
}
