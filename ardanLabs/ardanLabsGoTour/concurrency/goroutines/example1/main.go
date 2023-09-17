package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	Goroutines are functions that are created and scheduled to be run independently by the
	Go scheduler. The Go scheduler is responsible for the management and execution of
	goroutines.

	Let's first look at the scheduler semantics. When a Go program starts up, the Go runtime
	asks the machine (virtual or physical) how many operating system threads can run in
	parallel. This is based on the number of cores that are avaiable to the program. For each
	thread that can be run in parallel, the runtim will create an operating system thread (M) and
	attach that to a data structure that reps a logical processor (P) inside the program. This
	P and M rep the compute power or execution context for a running Go program.

	Also, an initial Goroutine (G) is created to manage the execution of instructions on a
	selected M/P. Just like an M manages the execution of instructions on the hardware, a G
	manages the execution of instructions on the M. This creates a new layer of abstraction
	above the operating system, but it moves execution control to the application level.

	Since the Go scheduler sits on top of the operating system scheduler, it's important to
	have some semantic understanding of the operating system scheduler and the constraints it
	applies to the Go scheduler and applications.

	The operating system scheduler has the job of creating the illusions that multiple pieces
	of work are being executed at the same time. Even when this is physically impossible. This
	requires some tradeoffs in the design of the scheduler. Before we go any further though,
	let's define some terms we'll see later on:

		Work: A set of instructions to be executed for a running application. This is accomplished
		      by threads and an application can have 1 to many threads.

		Thread: A path of execution that is scheduled and performed. Threads are responsible
		        for the execution of instructions on the hardware.

		Thread States: A thread can be in one of 3 states: Running, Runnable, or Waiting.
					   - Running means the thread is executing its assigned instructions (work)
					     on the hardware by having a G placed on the M.
					   - Runnable means the thread wants time on the hardware to execute its
					     assigned instructions and is sitting in a run queue.
					   - Waiting means the thread is waiting for something before it can resume
					     its work. Waiting threads are not a concern of the scheduler as the time
					     they are waiting depends on the works its doing and latencys that are
					     impacting it.

		Concurrency: This means undefined out of order execution. In other words, given a set
		             of instructions (work) that would be executed in the order provided, they
		             are executed in a different undefined order, but all executed. The key is
		             the result of executing the full set of instructions in any undefined
		             order produces the same result. We will say work can be done concurrently
		             when the order the work is executed in doesn't matter, as long as ALL the
		             work is completed.

		Parallelism: This means doing a lot of things at once. for this to be an option, we
					 need the ability to physically execute two or more operating system threads
					 at the same time on the hardware. O sea, we need multiple cores.

		CPU Bound Work: This is work (instructions) that does not cause the thread to naturally
						move into a WAITING state. Calculating fibonacci  numbers would be
						considered CPU-Bound work

		I/O Bound Work: This is work that does cause the thread to naturally move into a
						WAITING state. Fetching data from different URLs would be considered
						I/O-Bound Work.

		Synchronization: When two or more Goroutines will need to access the same memory location
						 potentially at the same time, they need to be synchronized and take
						 turns. If this synchronization doesn't take place, and at least one
						 Goroutine is performing a write, we can end up with a data race. Data
						 races are a cause of data corruption bugs that can be nasty to find.

		Orchestration: When two or more Goroutines need to signal each other, with or without
					   data, orchestration is the mechanic required. If orchestration doesn't
					   take place, guarantees about concurrent work being performed and
					   completed will be missed. Thsi can cause all sorts of data corruption
					   bugs.

		There are lots of little details related to the scheduling semantics, so to review
		more of those notes I can go to the links at the bottom of this lesson in Extra Reading
		"Scheduling in Go" parts 1-3

*/

// Now let's look at concurrency basics. We'll start with a basic problem that needs some
// orchestration, o sea, 2 or more goroutines signaling each other to know when or what
// work to handle

func init() {
	// The call to GOMAXPROCS is used to the  Go program in a single thread. This forces
	// our machine to use only one thread, regardless of how many cores it has and thus our
	// Go program will be single threaded and have a single P/M to execute all Goroutines.
	// The function is capitalized because it's also an environment var. The call to the
	// function overwrites the variable.

	// Allocate the logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// This function is an important function when you set CPU quoteas to a container config
	// for example. When passing 0, the number of threads the Go program will be using is
	// reported. We must make sure that number matches the number of OS threads we have
	// available in our containerized environment. If the numbers are not the same, the Go
	// program won't run as well as it otherwise could. You might want to use the environment
	// variable or this call to match things up.
	fmt.Println(runtime.GOMAXPROCS(0))
}

func main() {

	// In this program we hve to solve an orchestration problem. The main goroutine can't
	// allow the main function to return until there is a guarantee to two Goroutines being
	// created finish their work first. A WaitGroup is a perfect tool for orchestration
	// problems that don't require data to be passed between Goroutines. The signaling here
	// is performed thorugh an API that allows a Goroutine to wait for other Goroutines to
	// that they are done.

	// In this code, a WaitGroup is constructed to its zero value state and then immediately
	// the Add method is called to set the WaitGroup to 2 which will match the number of
	// Goroutines created.
	var wg sync.WaitGroup
	wg.Add(2)
	// When you know how many Goroutines upfront that will be created, you should call Add once
	// with that number. When you don't know (like in a streaming service) then calling Add(1)
	// is acceptable and reasonable.

	fmt.Println("Start Goroutines")

	// create a goroutine from the  lowercase function
	go func() { // NOTE the goroutine is started with 'go' then the anony func
		lowercase()
		wg.Done() // NOTE here is where the goroutine is marked off the WaitGroup list as done
	}()

	// create a goroutine from the upper case function
	go func() {
		uppercase()
		// wg.Done() // <-- If we forget to call Done then we'll get a DEADLOCK since the WaitGroup will
		// never get to 0 and be blocked forever
	}()
	// Literal functions (anony funcs) are declared and executed with the use of the keyword 'go'
	// At this point, you are telling Go scheduler to execute these functions concurrently. To
	// execute them in an undefined order. Inside the implementation of each Goroutine is the
	// call to Done. That call is what decrements the WaitGroup by 1. Once both calls to Done
	// are made, the WaitGroup will change from 2 to 0, and then the main GoRoutine will be
	// allowed to continue, be unblocked from the call to Wait, finishing the program.

	// at the end of main is the call to Wait, o sea, "wait for the goroutines here. Wait holds"
	// Wait holds the main Goroutine from causing the function to return. When the main function
	// returns, the Go program is shut down with extreme prejudice. This is why mananging the
	// orchestration with the proper guarantees is so important. The Wait call will BLOCK
	// until the WaitGroup is set back to 0
	fmt.Println("Waiting to finish")
	wg.Wait() // <-- if we comment out then there's no guarantee that the program will allow the grs
	// to finish. It might, but there's no guarantee. Since the calls to Println are sys calls and
	// thus allow the scheduler to make a context switch, delaying the runtime, one of the grs
	// could actually complete, but maybe not

	fmt.Println("\nTerminating Program")
}

// lowercase displays the set of lowercase letters three times\
func lowercase() {
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

// uppercase displays the set of uppercase letters 3 times
func uppercase() {
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

// When we build and run this program, we'll see how it runs concurrently. The second Goroutine
// created was scheduled first. It got to finish its work and then the other goroutine ran.
// Both ran to completion before the program finished. The next time we run this program,
// there is no guarantee we'll see the same output. The only guarantee is that 'main' won't
// finish until the two grs are complete.

// we could run this program 100 times and see the same output, but there is no guarantee it
// will happen again. It may be highly probable, but not guaranteed. Especially not guaranteed
// across different versions, OSs and architectures.
