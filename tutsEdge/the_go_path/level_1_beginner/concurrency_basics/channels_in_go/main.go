package main

import (
	"fmt"
	"math/rand"
)

// At some point you'll need to communicate between you goroutines. We do this with channels
// which are basically pipelines to communicate between routines.

func CalculateValue(values chan int) { // note that we aren't returning anything but we are...
	value := rand.Intn(10)
	fmt.Printf("Value Calculated: %d\n", value)
	//...passing something back to the main routine through the chan which was an arg of the func
	values <- value
}

func main() {
	values := make(chan int)
	// so after creating the int channel with 'make' we pass it to the function
	// it then does what it needs to do on the goroutine and passes it back via the chan
	go CalculateValue(values)

	value := <-values // this is a blocking call. So the program won't close until something
	// is received.
	fmt.Println(value)

}
