package main

import (
	"fmt"
)

// Data races are one of the most common and hardest things to debug in concurrent
// coding. This is when two goroutines access the same var concurrently and at least
// one of them is a write. Meaning that since we don't control which routine finishes
// first, we don't necesarily know what the value of that var will be when both grs
// complete. The GRs are 'racing' to get to the var first.

func main() {

	// Here is an example of a data race:
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // first conflicting access
		c <- true
	}()
	m["2"] = "b" // second conflicting access
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// we can diagnose such bugs with the built-in data race detector. Just add -race to
// the go command when running or building. If a data race if found then Go creates
// a report with a stack trace as well as where the the channels were created

// above we are adding in some examples of other data races that are common, so we 
// can avoid these patterns, but in the end we have the data race checker to help
// us identify these bugs.


