/*
	We will use 'go build -race' or 'go run main.go -race' with these

	A data race is when two or more GoRoutines are trying to access the same
	memory location at the same time WHERE AT LEAST ONE OF THEM IS PERFORMING
	A WRITE. When this happens, it is impossible to predict the result. These
	types of bugs are difficult to find because they cause issues that always
	appear random, Nasty.

	In addition to this we have the CPU Cacheline problem with "false sharing"
	Scott Meyers, CPU Caches and Why You Care:  https://youtu.be/WDIkqP4JbkE?t=1809

	First let's look at a Data Race Example:
	Below is a great example of a data race and how they can be hidden for years
	and eventually show up at odd times and cause data corruption
*/

package main

import (
	"fmt"
	"log"
	"sync"
)

// counter is a ver increment by all goroutines. Right off the bat this will be our
// problem because we will have multiple goroutines accessing this shared global memory
var counter int

func main() {

	// number of goroutines we will use
	const grs = 2

	// wg is used to manage concurrency
	var wg sync.WaitGroup
	wg.Add(grs) // NOTE here we use a const instead of a static number

	// Let's create the 2 goroutines. NOTE the different way we are creating goroutines here
	// instead of what we've done up until now
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {
				// Capture the value of Counter - Read Operation
				value := counter

				// Increment our local value of Counter
				value++

				// adding some logging
				log.Println("logging")

				// Store the value back into Counter - Write Operation
				counter = value
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// Now we get the right answer here, but just cuz we get the right answer doesn't mean that
// everything is kosher. Our operations are happening uninterrupted and that's why we get
// the correct answer. But what if we put some logging in as we added above. With the logging
// we don't always get 4, sometimes we get 3 and sometimes 2. Why?? When we run the data we
// run into a data race bug that always existed byt wasn't happening. The call to log is
// now causing the scheduler to make a context switch between the two grs at a bad time. So
// after the modify op a context switch is taking place. The three operations are no longer
// uniterrupted and gr2 ends up with its local value being wrong by the time it completes
// the write op. That's why we see it happening 'randomly'. But for this reason Go has its
// Race Detector. 
// 'go build -race' 

// Running this and then running the file we can see there is a data race and its detected
// with our without the log statement inserted. When a data race is detected, the program
// panics and provides a trace. The trace shows where there was unsynchronized access to 
// the same shared state where at least  one access was a write. 

// We can clearly see the unsynchronized read and write. As a side note, the ++ operation
// would also be a data race if the code was accessing the counter variable directly. The ++
// operation is a read,modify and write operation underneath (and might be the reason why
// my program is actually saying there are two data races) and the OS could easily context
// switch in the middle of that. 

// To fix this we have two tools. 
// - atomic instructions
// - mutexes

// First lets look at atomic instructions in the next example.

