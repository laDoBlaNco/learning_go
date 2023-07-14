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

var pl = fmt.Println

func main() {

	// for - initialization; condition; postStatement++ {BODY]
	for x := 5; x >= 1; x-- {
		pl(x)
	}
	// the for loops have a similar scope as functions. That x doesn't exist outside

	// While Loops with Go 4
	fx := 1
	for fx < 5 {
		pl(fx)
		fx++
	}

	nums := []int{1, 2, 3}
	for _, num := range nums {
		pl(num)
	}
	val := 1
	for true {
		if val == 5 {
			break
		}
		pl(val)
		val++
	}

	// so for the challenge we need to create a random number and then get user
	// input on what their guess of that number is and using conditionals respond
	// "higher" or "lower" until they get the number.

	// The guessing will need to be in a loop (either while (for) or for{} )
	seedVal := time.Now().Unix()
	rand.Seed(seedVal)
	randNum := rand.Intn(50) + 1
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Guess a number between 0 and 50: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	for {
		if guess > randNum {
			pl("Guess something lower!")
		} else if guess < randNum {
			pl("Guess something higher!")
		} else {
			pl("You got it!\nNice guess!")
			break
		}
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err = strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Challenge: Create a random game
// Guess a number between 0 and 50:
// Random number is: 25
// 20
// Higher
// 40
// Lower
