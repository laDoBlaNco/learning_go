package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
	"time"
)

// ... for quick debugging
var p = fmt.Println

// Go by Example: Channel Synchronization
// We can use channels to syncronize execution across goroutines. Here's an
// example of using a blocking receive to wait for a goroutine to finish. When
// waiting for multiple goroutines to finsh, you may prefer to use a WaitGroup

// This is the function we'll run in a goroutine. The 'done' channel will be
// used to notify another goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// send a value to notify that  we're done.
	done <- true
}

func main() {

	// Let's start a worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// This will block until we receive a notification from the worker on the
	// channel
	<-done
	// without this <- done waiting for an blocking the program from closing, it
	// would close before the worker even started. So the blocking is happening here
	// with the WAITING for the channel, not the work that's being done in the
	// goroutine

}

// Next let's look at Channel Directions
