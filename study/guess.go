// This guessing program is player guess the random number in 10 chances

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it? \n You have 10 chance to guess. Let's start!")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("You have ", 10-guesses, " chance left.")
		fmt.Print("Make a guess (between 1 to 100): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if target > guess {
			fmt.Println("Oops, Your guess was LOW!")
		} else if target < guess {
			fmt.Println("Oops, Your guess was HIGH!")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!!")
			break
		}

		if !success {
			fmt.Println("Sorry, You didn't guess my number. It was [", target, "].")
		}

	}

}
