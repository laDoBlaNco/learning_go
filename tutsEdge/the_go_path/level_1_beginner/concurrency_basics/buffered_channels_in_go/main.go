package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	for i := 0; i < 10; i++ { // sending a value to values channel 10 times
		value := rand.Intn(10)
		fmt.Printf("Value Calculated: %d\n", value)
		values <- value
	}
}

func main() {
	// A buffered channel allows us to block only when our channel is full. Its basically
	// the capacity of our channel to be able to hold more than 1 item a time. So it doesn't
	// block on every single send. We do that by just adding one more arg to 'make'
	values := make(chan int, 2) // here we added a limit of 10 to make this a 10 buffer chan
	go CalculateValue(values)

	for i := 0; i < 10; i++ { // receive a value to values channel 10 times
		time.Sleep(1 * time.Second)
		value := <-values
		fmt.Println(value)
	}
}

// This is useful because it allows us to minimize the amount of tasks tha can processed
// at the same time. So really it helps with the balancing of the work between creation and
// consumption.
