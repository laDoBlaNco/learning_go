package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println // I didn't know you could do that

// line comment

/*
Block comments
*/

func main() {
	pl("Hello gophers!") // didn't know you could do this at all. Alreaady learned
	// something new and I just started. May not be very idiomatic Go though.
	fmt.Println("Hello Gophers!")

	fmt.Print("What is your name? ")
	reader := bufio.NewReader(os.Stdin)  // using bufio to get user input from cli
	name, err := reader.ReadString('\n') // bufio.ReadString until user hits \n

	if err != nil {
		log.Fatal(err)
	}
	pl("Hello", name)

}
