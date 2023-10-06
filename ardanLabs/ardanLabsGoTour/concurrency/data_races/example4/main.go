package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
	Read/Write mutexes

	There is a second type of mutex called a read/write mutex. It allows us to separate
	the locks around reads an writes. This is important since reading data doesn't
	pose a threat unless a Goroutine is attempting to write at the same time. So this
	type of mutex allows multiple Goroutines to read the same memory at the same time
	So this type of mutex allows multiple grs to read the same memory at the same time.
	As soon as a write lock is requested, the reads are no longer issued, the write takes
	place, the reads will then start again

*/

// First let's have a data slice that will be shared
var data []string

// then a rwmutex to define the critical section of code
var rwMutex sync.RWMutex

// number of reads occuring at any given time
var readCount int64

// then let's initiate our func
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// wg will be used to manage concurrency as before
	var wg sync.WaitGroup
	wg.Add(1)

	// create our writer gr
	go func() {
		for i := 0; i < 10; i++ {
			writer(i)
		}
		wg.Done()
	}()

	// create either other reader grs
	for i := 0; i < 8; i++ {
		go func(id int) {
			for {
				reader(id)
			}
		}(i)
	}

	// Wait for the write gr to finish
	wg.Wait()
	fmt.Println("Program complete")
	fmt.Println(data) 
}

// Let's create  our helper functions, first writer which adds a new string to the slice
// at random intervals
func writer(i int) {
	// only allow one gr to read/write to the slice at a time.
	rwMutex.Lock()
	{
		// capture the current read count
		// keep this safe though we can do without this call
		rc := atomic.LoadInt64(&readCount)
		
		// perform some work since we have a full lock
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) 
		fmt.Printf("****> : Performing Write : RCount[%d]\n",rc) 
		data = append(data,fmt.Sprintf("String: %d",i)) 
	}
	rwMutex.Unlock() // release the lock
}

// now we have 'reader' which wakes up and iterates over the slice
func reader(id int){
	
	// any gr can read when no write operation is taking place
	rwMutex.RLock() 
	{
		// Increment the read count value by 1
		rc := atomic.AddInt64(&readCount,1) 
		
		// Perform some read work and display values
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) 
		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n",id,len(data),rc) 
		
		// decrement the read count value by 1
		atomic.AddInt64(&readCount,-1) 
	}
	rwMutex.RUnlock() 
	// Release the read lock  
}
// NOTE the use of Lock and RLock as well as Unlock and RUnlock. If we get these mixed p
// we will have major issues. 

/*
	NOTE:
	- Goroutines need to be coordinated and synchronized
	- When two or more goroutines attempt to access the same resource, and one of them is a write
	  operation, we have a data race
	- Atomic functions and mutexes can provide the support we need. 
	

*/
