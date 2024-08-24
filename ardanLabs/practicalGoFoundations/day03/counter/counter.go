package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// In some situations we will have data races. This is when we have a bunch of goroutines all
// making changes on the same data, for example our 10 counters counting to 10,000. which
// should give us 100_000 in the end, but doesn't due to the numbers being cached and adjust are
// all different, so there's no guarantee that using 10 goroutines to work on the same counter
// will work as you expect, which adding something to manage it, like sync.Mutex. Now with the
// mutex we get 100_000 and no race conditions

func main() {
	/* Solution 1: using mutex
	var mu sync.Mutex // its customary to put the mutex var on top of the item or struct they are
	// guarding.  Here we are guarding the 'count'
	count := 0
	*/

	// Solution 2: using sync/atomic
	// count := int64(0)
	var count int64
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				/*
					mu.Lock() // now we can put the candado on and hand the keys to our goroutines
					count++
					mu.Unlock()
				*/
				atomic.AddInt64(&count, 1) // now since we are at the atomic level, there's no need
				// for caching or mutex because the original counter is being changed everytime by
				// all of the goroutines.But think carefully if we really need it. This is for
				// squeezing all performance from our code. If we are using it a lot, we are using
				// it wrong more than likely. Preference for a gopher:
				// 1. Channels
				// 2. Sync
				// 3. sync/atomic

				// With concurrency problems there are mainly 2 things to look at.
				// 1. is it a problem of syncronization, meaning a bunch of routines working on the
				// 	  shared resource
				// 2. then there's orchestration, which is where we use WaitGroups.
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// We also have sync.Atomic. So syncing at a much lower level. Allowing us to do things at
// the atomic level so there is no copying or caching.

// Next we'll look at Selects after the break ;)

