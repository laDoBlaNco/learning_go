package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for state in Go is communication over channels.
// We saw this for example with 'worker pools'. There are a few other options
// for managing state though. Here we'll look at using the sync/atomic pkg
// for atomic counters accessed by multiple goroutines

func main() {
	// we'll use an unsigned int to represent our  (always positive) counter
	var ops uint64

	// a WaitGroup will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// we'll start 50 goroutines that each incremenet the counter exactly 100 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// To atomically increment the counter we use AddUint64, giving it the
		// memory address of our ops counter with the & pointer syntax
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	// Wait until all the goroutines are done
	wg.Wait()

	// Its safe to access ops now because we know no other goroutine is writing to
	// it. Reading atomics safely while they are being updated is also possible,
	// using functions like atomic.LoadUint64
	fmt.Println("ops:", ops)
}

// We expect to get exactly 50,000 operations. Had we used the non-atomic ops++ to
// increment the counter, we'd likely get a different number, changing between runs
// because the goroutines would interfere with each other. Moreover, we'd get data
// race failures when running with the -race flag.

// Next we'll look at mutexes, another tool for managing state.
