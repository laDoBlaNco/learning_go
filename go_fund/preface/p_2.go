package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// create a channel of type int
	ch := make(chan int)

	// create a wait group to track the goroutine
	var wg sync.WaitGroup

	// create some goroutines to listen for message on the channel
	for i := 0; i < 4; i++ {
		// increment the wait group
		wg.Add(1)

		// create a goroutine to call the sayHello function
		go func(id int) {
			// call the sayHello function
			sayHello(id, ch)

			// decrement the wait group when the sayHello function exits
			wg.Done()
		}(i + 1)
	}

	// send messages on the channel
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// close the channel to signal the goroutine to exit
	close(ch)

	// wait for all goroutines to finish
	wg.Wait()
}

func sayHello(id int, ch chan int) {
	// listen for messages on the channel
	// loop exits when the cahnnel is closed
	for i := range ch {
		fmt.Printf("Hello %d from goroutine %d\n", i, id)

		// simulate a long-running task
		time.Sleep(500)
	}
}
