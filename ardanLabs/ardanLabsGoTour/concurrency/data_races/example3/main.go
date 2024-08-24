// This program will show us how to use a mutex to define critical sections of code
// that need  synchronous access

package main

import (
	"fmt"
	"sync"
)

// let's start with our global counter for the shared state
var counter int // NOTE since we aren't using atomics no need for the precision

// mutex is used to define critical sections of code
// So if we want to keep the 3 lines of code we had before then atomics aren't going to
// work. What we use is a mutex. A mutex lets us box a group of code so only one gr at a
// time can execute that code.
var mutex sync.Mutex

func main() {

	// Number of grs to use
	const grs = 2

	// wg is used to manage concurrency
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {

				// Only allow one goroutine thorugh this critical section at a time.
				mutex.Lock()
				{
					// Capture the value of counter
					value := counter

					// increment our local value of counter
					value++

					// Store the value back into counter
					counter = value
				}
				mutex.Unlock()
				// Release the lock and allow any waiting goroutine through
			}
			wg.Done()
		}()
	}

	// Wait for the goroutines to finish
	fmt.Printf("Final Counter: %d\n", counter)
}

// with this code in place, the scheduler will only allow one Goroutine to enter the code
// block at a time. Its important to understand that a mutex is not a queue. The first GR
// that calls Lock isn't necessarily the first GR who gets the Lock. There is a fairness
// based algorithm but this is done on purpose so people don't use mutexes as queues.

// It's important to remember the  Lock creates back pressure, so the longer it takes to get
// from Lock to Unlock, the more chance of Goroutines  waiting for their turn. If you forget
// to Unlock, then all Goroutines waiting will deadlock. This is why it's critical that the
// call to lock and Unlock happen in the same function. Make sure we are doing the bare
// minimum synchronization you need in a code block, but at least the minimum.

// note that this results in an incorrect result. Its very bad code where we are trying to
// get in and out of the lock so quickly, that we lose our synchronization and the race detector
// can't even catch it.
