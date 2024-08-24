package main

import (
	"fmt"
	"sync"
)

// Go by Example: Mutexes
// In the previous example we saw how to manage simple counter state using
// atomic operations. for more complex state we can use a mutex to safely
// access data cross multiple goroutines

// Container holds a map of counters; since we want to update it concurrently
// from multiple goroutines, we add a Mutex to synchronize access. Note that
// mutexes must not be copied, so if this struct is passed around, it should be
// done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Lock the mutex before accessing counters; unlock it at the  end of the function
// using a defer statement.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// Note that the zero value of a mutex is usable as-is, so no initiialization is
// required here
func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Run severlal goroutines concurrently; note that they all access the same
	// Container, and two of them access the same counter.
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}

// Running the program shows that the counters updated as expected
// Next we'll look at implementing this same state management t ask using
// only goroutines and channels
