package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

/*

	Preemptive Scheduler -

	Even though the scheduler runs within the scope of the application, its important
	to see how the schedule is preemptive. This means you can't predict when a context
	switch will take place and this will change every time you run the program.

*/

// In this example we see how the gr scheduler will time slice grs on a single thread

func init() {
	// allocate one logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.GOMAXPROCS(0))
}

func main() {

	// wg is used to manage concurrency
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")

	// Create the first gr and manage its lifecycle here.
	go func() {
		printHashes("A")
		wg.Done()
	}()

	// Create the second gr and manage its lifecycle here
	go func() {
		printHashes("B")
		wg.Done()
	}()

	// Wait for the grs to finish
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printHashes calculates the sha1 hash for a range of numbers and prints each in
// hex encoding
func printHashes(prefix string) {

	// print each hash from 1 to 10. Change this to 50000 to see how the scheduler behaves
	for i := 1; i <= 50000; i++ {

		// convert i to string
		num := strconv.Itoa(i)

		// calculate hash for the string num
		sum := sha1.Sum([]byte(num))
		
		// Print prefix: 5-digit-number: hex encoded hash
		fmt.Printf("%s: %05d: %x\n",prefix,i,sum) 
	}
	
	fmt.Println("Completed",prefix) 

}
// The function is performing a lot of io bound work that has the potential of being context
// switched and therefore with only 10, we get the same results typically, but with 50000 the results
// are different everytime

/*
	NOTE:
	1. Goroutines are functions that are scheduled to run independently
	2. We must always maintain an account of running goroutines and shut down cleanly
	3. Concurrency is not parallelism
	4. Concurrency is about DEALING with lots of things at once
	5. Parallelism is about DOING lots of things at once. 

	"Parallelism is about physically doing two or more things at the same time. Concurrency
	is about undefined, out of order, execution." - William Kennedy
	
	"By default, goroutines shouldn't outlive the function they were created from, this forces
	us into extremely good design posture." - Peter Bourgon


*/
