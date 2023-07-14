package main

import (
	"fmt"
	"sync"
	"time"
)

// GBE: WaitGroups

// To wait for multiple goroutines to finish, an alternative to blocking the program
// from ending is using waitgroup.

// first we have a function that we'll run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// This WaitGroup is used to wait for all the goroutines launched here to
	// finish. Note: if a WaitGroup is explicitly passed in functions, it should
	// be done BY POINTER. I assume because if not then we'll have our WaitGroup and
	// a COPY of our WaitGroup
	var wg sync.WaitGroup

	// Then we launch several goroutines and increment the WaitGroup counter for each
	// Note that we have to tell the WaitGroup how many routines to wait for, so I can
	// think of it as someone sitting with a clipboard and managing how many jobs start
	// and checking them off after the each end
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// avoid re-use of the same i value in each goroutine closure.
		// Note: this is what Jayson was talking about in ZTM. With loops in concurrency
		// we need to use copies rather than the actual iteration var
		i := i // so inside the loop i shadows i, but they are differnt more details
		// below taken from the Go FAQ ;)

		// we then wrap the worker call in a closure that makes sure to tell the wg
		// that this worker is done. This way the worker itself does not have to be
		// aware of the concurrency primitives involved in its execution
		go func() {
			defer wg.Done() // telling the wg we are done after each worker finishes
			worker(i)
		}()
	}
	// this will then block (our guy with the clipboard) until the wg counter goes back
	// to 0(workers getting checked off); all the workers notified they're done
	wg.Wait()
}

// Note: this approach has no straightforward way to propagate errors from workers. for
// more advanced use cases, consider using the 'errgroup package'

// Now let's check out 'Rate Limiting' next

/*
What happens with closures running as goroutines?
Some confusion may arise when using closures with concurrency. Consider the following
program:

func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting
    for _ = range values {
        <-done
    }
}
One might mistakenly expect to see a, b, c as the output. What you'll probably see instead
is c, c, c. This is because each iteration of the loop uses the same instance of the
variable v, so each closure shares that single variable. When the closure runs, it
prints the value of v at the time fmt.Println is executed, but v may have been modified
since the goroutine was launched. To help detect this and other problems before they happen,
run go vet.

To bind the current value of v to each closure as it is launched, one must modify the inner
loop to create a new variable each iteration. One way is to pass the variable as an
argument to the closure:

    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
In this example, the value of v is passed as an argument to the anonymous function.
That value is then accessible inside the function as the variable u.

Even EASIER is just to create a new variable, using a declaration style that may seem
odd but works fine in Go:

    for _, v := range values {
        v := v // create a new 'v'.
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }
This behavior of the language, not defining a new variable for each iteration, may have
been a mistake in retrospect. It may be addressed in a later version but, for
compatibility, cannot change in Go version 1.

*/
