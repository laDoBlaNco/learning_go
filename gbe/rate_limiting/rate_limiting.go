package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and
// maintaining quality of service. Go elegantly supports rate limiting with
// goroutines, channels, and tickers.

func main() {

	// First we'll look at basic rate limiting. Suppose we want to limit our
	// handling of incoming requests. We'll serve these requests off a channel
	// of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds. This
	// is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each
	// request, we LIMIT ourselves to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter                               // this blocks until a message is sent every 200 milliseconds
		fmt.Println("request", req, time.Now()) // thus stopping this from happening
		// until after the 200 milliseconds passes on each iteration
	}

	// We may want to allow short bursts of requests in our rate limiiting scheme
	// while preserving the overall rate limit. We can accomplish this by buffering
	// our limiter channel. This burstyLimiter channel will allow bursts of up to
	// 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds will try to add a new value to burstyLimiter, up to its
	// limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// now simulate 5 more incoming requests. The first 3 of these will benefit
	// from the burst capability fo the burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now()) 
	}
}

// running the program we see the first batch of requests once every ~200
// milliseconds as desired. For the second batch of requests we serve first 3 
// immediately becasue of the burstable rate limiting, then serve the remaining
// 2 with ~200ms delays each.
