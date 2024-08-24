package main

/*
	This example shows how to use the atomic package to provide safe access to NUMERIC 
	types with shared memory. 
	
	Atomics provide synchronization at the hardware level. Because of this, it's limited
	to words and half-words of data. So they're great for counters or fast switching 
	mechanics. The WaitGroup APIs use atomics actually. 
	
	So what changes do we need to make to apply atomics to the code?

*/

import(
	"fmt"
	"sync"
	"sync/atomic" // atomic is part of the same sync pkg
)

// here is our global counter
var counter int64 // right off the bat notice that we need to be specific about precision
// when using atomics. So no more 'int' now its 'int64'. The atomic functions only work with
// precision based integers. 

func main(){
	
	// number of grs
	const grs = 2
	
	// wg to manage concurrency
	var wg sync.WaitGroup 
	wg.Add(grs) 
	
	
	// we have our grs
	for g:=0;g<grs;g++{
		go func(){
			for i:=0;i<2;i++{
			// next we remove the manual read, modify, and write code for a call to atomic
			// AddInt64. This one call handles all the code we replaced. All functions related
			// with atomiic package take the address to the shared state to be syncrhonized
			// Synchronization only happens at the address level. So different grs calling the
			// same function, but at different addresses, won't be synchronized.
				atomic.AddInt64(&counter,1) 
			}
			wg.Done()
		}()
	}
	
	wg.Wait()
	fmt.Println("Final Counter:",counter) 
}
// and with that, the data race is gone.  As we can see our first parameter to the atomic
// is always the address to a precision based integer or pointer. There is also a type
// named Value that provides a synchronous value with a small API

// The second way to address data races is with mutexes. Let's take a look at that in the next 
// example.


