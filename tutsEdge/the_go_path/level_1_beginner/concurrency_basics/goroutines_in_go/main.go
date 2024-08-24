package main

import (
	"fmt"
	"time"
)

// The standard practice is to use a actual named function to start in a goroutine
// This doesn't block the main routine. We can see that as the main routin continues to print
// while we wait on this goroutine.
func HelloWorld(name string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Hello, %s\n", name)

}

func main() {
	// there a couple of options to  use goroutines. This first one is using an anony func
	go func() {
		fmt.Println("Hello world!")
	}()

	// here's our second way of running a goroutine, with an actual function
	go HelloWorld("Ladoblanco")

	fmt.Println("I should be printed first")

	time.Sleep(2 * time.Second) // this is a super hacky way to get the goroutine to run.
	// basically blocking the main thread to give us time to let the goroutine run.
}
