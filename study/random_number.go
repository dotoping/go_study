// Shows random numbers as much as a range of numbers.

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
	var rangeNumber int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What range of random number do you want to get? : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	rangeNumber, err = strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("How many random number do you want to see? : ")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	countInput, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	for countNumber := 0; countNumber < countInput; countNumber++ {
		randNumber := rand.Intn(rangeNumber) + 1
		fmt.Println(randNumber)
	}

}
