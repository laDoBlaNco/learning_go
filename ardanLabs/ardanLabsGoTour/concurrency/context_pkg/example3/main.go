package main

import (
	"context"
	"fmt"
	"time"
)

/*
	In this program we see the use of WithDeadline


*/

type data struct {
	UserID string
}

func main() {

	// set the deadline
	deadline := time.Now().Add(150 * time.Millisecond)

	// create a context that is both manually cancellable (like WithCancel) and will signal a
	// cancel at the specified date/time (deadline)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// let's create a channel to recieve a signal that work is done
	ch := make(chan data, 1)

	// ask the goroutine to do some work for us
	go func() {
		// simulate work
		time.Sleep(200 * time.Millisecond)

		// report the work is done
		ch <- data{"123"}
	}()

	// wait for the work to finish. If it takes too long move on
	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}

// now let's see WithTimeout

