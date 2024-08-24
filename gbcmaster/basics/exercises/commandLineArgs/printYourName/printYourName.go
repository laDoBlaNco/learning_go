package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

func main() {
	fmt.Println(os.Args[1])

	fmt.Println("Hi", os.Args[1])
	fmt.Println("How are you?")

}

// on my first version I used a var, but apparently he wanted it directly from
// os.Args
