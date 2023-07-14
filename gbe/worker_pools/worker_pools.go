package main

import (
	"fmt"
	"time"
)

// GBE: Worker Pools
// In this example we are going to look at how to implement a worker pool
// using goroutines and channels

// Here's the worker, of which we'll run several concurrent instances. These
// workers will receive work on the jobs channel and send the corresponding
// results on a results channel. We'll sleep a second per job to simulate
// an expensive task
func worker(id int, jobs <-chan int, results chan<- int) { // (id, sender, reciever)
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// In order to use our pool of workers, we need to send them work and collect
	// their results. We make 2 channels for this.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, intially blocked because there are no jobs  yet
	// so basically we are starting our workers and they are waiting for work
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results) // (id,sender channel,receiver channel)
	}

	// Here we are sending 5 jobs and then closing that channel to indicate that
	// there's no more work.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work. This also ensures that the
	// worker goroutines have finished. An alternative way to wait for multiple
	// goroutines is to use a 'waitgroup'
	for a := 1; a <= numJobs; a++ {
		<-results // this will block the program from ending till everything is collected
	}

}

// our running program should show 5 jobs being executed by various workers. The
// program only takes about 2 seconds despite doing around 5 seconds of total work
// because there are 3 workers opeprating concurrently

// program runs a little different as I play with the number of workers. First 5 and then
// only 2.
