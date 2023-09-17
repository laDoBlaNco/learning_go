// This sample program shows howwe create goroutines and how the scheduler behaves
// with two contexts
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// let's first allocate two logical processors instead of the 1 we usually do. So this is now
	// a multi-threaded machine
	runtime.GOMAXPROCS(2)

	// fmt.Println(runtime.GOMAXPROCS(0))
}

func main() {

	// let's start with wg to wait for the program to finish. Adding in a count of 2
	// one for each gr
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Now we start with our grs, as usual using an anony func with 'go'
	go func() {
		// display the alphabt three times
		for count := 0; count < 3; count++ {
			for r := 'a'; r <= 'z'; r++ {
				fmt.Printf("%c ", r)
			}
		}

		// knowck this gr off our wait list
		wg.Done()
	}()
	
	// Now let's do our second gr the same way
	go func(){
		for count:=0;count<3;count++{
			for r:='A';r<='Z';r++{
				fmt.Printf("%c ",r) 
			}
		}
		
		wg.Done() 
	}()


	// and heres the actual waiting call
	fmt.Println("Waiting to finish") 
	wg.Wait() 
	
	fmt.Println("\nTerminating Program") 
}

// Note the results are different as here we have a real concurrent/parallel program

