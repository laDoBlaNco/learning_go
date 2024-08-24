package main

import (
	"fmt"
	"time"
)

// Timers  are for when you want to do something once in the future - 'tickers' are
// for when you want to do something repeatedly at regular interfals. Here's an
// example of a ticket that ticks periodically until we stop it.
func main() {

	// Tickers use a similar mechanism to timers: a chan that is sent values.
	// Here we'll use  the 'select' built-in on the chan to await the values
	// as they arrive every 500ms
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t) // t time.Time
			}
		}
	}()
	// Like timers, tickers can be stopped. Once a ticker is stopped it won't receive
	// any more values on its channel. We'll stop ours at 1600ms.
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

// I've change this program to run for 5 seconds with a tick every second, then it
// stops.

// Now let's take a look at Worker Pools???
