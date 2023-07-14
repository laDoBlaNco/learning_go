package main

import (
	"fmt"
	"log"

	"examples.com/greetings"
)

func main() {
	// Set props of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names:
	names := []string{"L@D0Bl@Nc0", "0D@lis", "K3l3n", "X@v13r"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	// Request a greeting message.
	// message,err := greetings.Hello("L@D0Bl@Nc0")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message (map of messages)
	// to the console.
	fmt.Println(messages)
}
